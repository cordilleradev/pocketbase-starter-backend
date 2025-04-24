package sep10

import (
	"strings"
	"time"

	"github.com/stellar/go/keypair"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
)

// ValidationResult represents the result of validating a SEP-10 challenge transaction
type ValidationResult struct {
	ClientAccountID   string
	ClientMemo        *uint64
	ClientMuxedID     *uint64
	ClientDomain      string
	MatchedHomeDomain string
}

// ValidateSep10Challenge validates a SEP-10 challenge transaction
func ValidateSep10Challenge(
	challengeTxBase64 string,
	serverAccountID string,
	homeDomain string,
	webAuthDomain string,
	networkPassphrase string,
) (*ValidationResult, error) {
	// Parse the transaction
	tx, err := txnbuild.TransactionFromXDR(challengeTxBase64)
	if err != nil {
		return nil, errors.Wrap(err, string(ErrInvalidTransaction))
	}

	txe, ok := tx.Transaction()
	if !ok {
		return nil, ErrInvalidTransaction
	}

	// Verify sequence number is 0
	if txe.SequenceNumber() != 0 {
		return nil, ErrInvalidSequenceNumber
	}

	// Verify source account is the server account
	if txe.SourceAccount().AccountID != serverAccountID {
		return nil, ErrInvalidTransaction
	}

	// Check time bounds
	tb := txe.Timebounds()
	if tb.MaxTime == 0 && tb.MinTime == 0 {
		return nil, ErrInvalidTimeBounds
	}

	currentTime := time.Now().UTC().Unix()
	if currentTime < txe.Timebounds().MinTime || currentTime > txe.Timebounds().MaxTime {
		return nil, ErrChallengeExpired
	}

	// Check operations
	if len(txe.Operations()) == 0 {
		return nil, ErrNoOperations
	}

	// Validate first operation
	firstOp, ok := txe.Operations()[0].(*txnbuild.ManageData)
	if !ok {
		return nil, ErrInvalidOperationType
	}

	// Check if client account is provided in the first operation
	clientAccountID := firstOp.SourceAccount
	if clientAccountID == "" {
		return nil, ErrInvalidOperationSource
	}

	// Validate auth operation key format
	if !strings.HasSuffix(firstOp.Name, " auth") {
		return nil, ErrInvalidFirstOperation
	}

	// Extract and validate home domain
	opHomeDomain := strings.TrimSuffix(firstOp.Name, " auth")
	if opHomeDomain != homeDomain {
		return nil, ErrInvalidHomeDomain
	}

	// Check for web_auth_domain in the operations
	var foundWebAuthDomain bool
	var clientDomain string
	for i := 1; i < len(txe.Operations()); i++ {
		op, ok := txe.Operations()[i].(*txnbuild.ManageData)
		if !ok {
			return nil, ErrInvalidOperationType
		}

		// Verify web_auth_domain operation
		if op.Name == "web_auth_domain" {
			if op.SourceAccount != serverAccountID {
				return nil, ErrInvalidOperationSource
			}

			webAuthValue := string(op.Value)
			if webAuthValue != webAuthDomain {
				return nil, ErrInvalidWebAuthDomain
			}

			foundWebAuthDomain = true
		}

		// Check for client_domain operation
		if op.Name == "client_domain" {
			clientDomain = string(op.Value)
		}
	}

	if !foundWebAuthDomain {
		return nil, ErrInvalidWebAuthDomain
	}

	// Verify transaction signatures
	err = verifyTransactionSignatures(txe, clientAccountID, serverAccountID, networkPassphrase)
	if err != nil {
		return nil, err
	}

	// Parse result
	result := &ValidationResult{
		ClientAccountID:   clientAccountID,
		MatchedHomeDomain: opHomeDomain,
		ClientDomain:      clientDomain,
	}

	// Handle muxed accounts or memo
	if isMuxedAccount(clientAccountID) {
		muxedAccount, err := xdr.AddressToMuxedAccount(clientAccountID)
		if err != nil {
			return nil, ErrInvalidClientAccountID
		}

		id := uint64(muxedAccount.Med25519.Id)
		result.ClientMuxedID = &id
	} else if txe.Memo() != nil {
		if memoID, ok := txe.Memo().(txnbuild.MemoID); ok {
			id := uint64(memoID)
			result.ClientMemo = &id
		}
	}

	return result, nil
}

// verifyTransactionSignatures validates the signatures on the transaction
func verifyTransactionSignatures(txe *txnbuild.Transaction, clientAccountID string, serverAccountID string, networkPassphrase string) error {
	// Verify transaction has signatures
	signatures := txe.Signatures()
	if len(signatures) < 2 {
		return ErrClientSignatureMissing
	}

	// Verify server signature
	serverKP, err := keypair.Parse(serverAccountID)
	if err != nil {
		return ErrInvalidServerAccountID
	}

	txHash, err := txe.Hash(networkPassphrase)
	if err != nil {
		return errors.Wrap(err, string(ErrInvalidTransaction))
	}

	var serverSigFound bool
	var clientSigs []xdr.DecoratedSignature

	// Separate server and client signatures
	for _, sig := range signatures {
		if verifySignature(serverKP, txHash, sig.Signature) {
			serverSigFound = true
		} else {
			clientSigs = append(clientSigs, sig)
		}
	}

	if !serverSigFound {
		return ErrServerSignatureMissing
	}

	// For non-muxed accounts, verify the client's signature(s)
	if !isMuxedAccount(clientAccountID) {
		clientKP, err := keypair.Parse(clientAccountID)
		if err != nil {
			return ErrInvalidClientAccountID
		}

		// At least one signature must be valid for the client
		var validClientSig bool
		for _, sig := range clientSigs {
			if verifySignature(clientKP, txHash, sig.Signature) {
				validClientSig = true
				break
			}
		}

		if !validClientSig {
			return ErrInvalidSignature
		}
	}

	return nil
}

// verifySignature checks if a signature is valid for a given hash and keypair
func verifySignature(kp keypair.KP, hash [32]byte, signature []byte) bool {
	if kp == nil {
		return false
	}

	err := kp.Verify(hash[:], signature)
	return err == nil
}

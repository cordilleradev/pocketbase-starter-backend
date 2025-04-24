package sep10

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
)

// ChallengeParams represents the parameters required to build a SEP-10 challenge transaction
type ChallengeParams struct {
	ServerSignerSecret     string
	ClientAccountID        string
	HomeDomain             string
	WebAuthDomain          string
	Memo                   *uint64
	ClientDomain           string
	ClientDomainSigningKey string
	Timeout                time.Duration
	NetworkPassphrase      string
}

// BuildSep10Challenge builds a SEP-10 compliant challenge transaction
func BuildSep10Challenge(params ChallengeParams) (string, error) {
	// Validate server signing key
	serverKP, err := keypair.ParseFull(params.ServerSignerSecret)
	if err != nil {
		return "", ErrInvalidServerAccountID
	}

	// Validate client account
	var clientAccountAddress string
	var muxedClientAccount *xdr.MuxedAccount
	if isMuxedAccount(params.ClientAccountID) {
		if params.Memo != nil {
			return "", ErrMemoAndMuxedAccount
		}

		var tmpMuxedClientAccount xdr.MuxedAccount
		tmpMuxedClientAccount, err = xdr.AddressToMuxedAccount(params.ClientAccountID)
		if err != nil {
			return "", ErrInvalidClientAccountID
		}
		muxedClientAccount = &tmpMuxedClientAccount
		clientAccountAddress = muxedClientAccount.Address()
	} else {
		_, err = keypair.Parse(params.ClientAccountID)
		if err != nil {
			return "", ErrInvalidClientAccountID
		}
		clientAccountAddress = params.ClientAccountID
	}
	// Validate home domain
	if len(params.HomeDomain) == 0 {
		return "", ErrInvalidHomeDomain
	}

	if len(params.HomeDomain) > MaxHomeDomainLength {
		return "", ErrHomeDomainTooLong
	}

	// Validate web auth domain
	if len(params.WebAuthDomain) == 0 {
		return "", ErrInvalidWebAuthDomain
	}

	if len(params.WebAuthDomain) > MaxWebAuthDomainLength {
		return "", ErrWebAuthDomainTooLong
	}

	// Ensure web auth domain is valid URL
	_, err = url.Parse(params.WebAuthDomain)
	if err != nil {
		return "", ErrInvalidWebAuthDomain
	}

	// Generate 48 bytes of random data
	randomBytes := make([]byte, 48)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", ErrRandomGeneration
	}
	randomB64 := base64.StdEncoding.EncodeToString(randomBytes)

	// Determine network passphrase if not provided
	networkPassphrase := params.NetworkPassphrase
	if networkPassphrase == "" {
		// Default to public network
		networkPassphrase = network.PublicNetworkPassphrase
	}

	// Set default timeout if not provided
	timeout := params.Timeout
	if timeout == 0 {
		timeout = 15 * time.Minute
	}

	// Create time bounds
	now := time.Now().UTC()
	timeBounds := txnbuild.NewTimebounds(now.Unix(), now.Add(timeout).Unix())

	// Create operations
	homeDomainAuthKey := fmt.Sprintf("%s auth", params.HomeDomain)

	operations := []txnbuild.Operation{
		&txnbuild.ManageData{
			SourceAccount: clientAccountAddress,
			Name:          homeDomainAuthKey,
			Value:         []byte(randomB64),
		},
		&txnbuild.ManageData{
			SourceAccount: serverKP.Address(),
			Name:          "web_auth_domain",
			Value:         []byte(params.WebAuthDomain),
		},
	}

	// Add client domain operation if provided
	if params.ClientDomain != "" {
		if len(params.ClientDomain) > MaxClientDomainLength {
			return "", ErrClientDomainTooLong
		}

		if params.ClientDomainSigningKey == "" {
			return "", ErrInvalidClientDomainKP
		}

		clientDomainKP, err := keypair.Parse(params.ClientDomainSigningKey)
		if err != nil {
			return "", ErrInvalidClientDomainKP
		}

		operations = append(operations, &txnbuild.ManageData{
			SourceAccount: clientDomainKP.Address(),
			Name:          "client_domain",
			Value:         []byte(params.ClientDomain),
		})
	}

	// Create transaction
	var txOptions txnbuild.TransactionParams
	if params.Memo != nil {
		txOptions = txnbuild.TransactionParams{
			SourceAccount: &txnbuild.SimpleAccount{AccountID: serverKP.Address()},
			Operations:    operations,
			BaseFee:       txnbuild.MinBaseFee,
			Memo:          txnbuild.MemoID(*params.Memo),
			Preconditions: txnbuild.Preconditions{
				TimeBounds: timeBounds,
			},
		}
	} else {
		txOptions = txnbuild.TransactionParams{
			SourceAccount: &txnbuild.SimpleAccount{AccountID: serverKP.Address()},
			Operations:    operations,
			BaseFee:       txnbuild.MinBaseFee,
			Preconditions: txnbuild.Preconditions{
				TimeBounds: timeBounds,
			},
		}
	}

	tx, err := txnbuild.NewTransaction(txOptions)
	if err != nil {
		return "", errors.Wrap(err, "failed to build transaction")
	}

	// Sign transaction with server key
	tx, err = tx.Sign(networkPassphrase, serverKP)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign transaction")
	}

	// Convert to base64
	txBytes, err := tx.Base64()
	if err != nil {
		return "", errors.Wrap(err, "failed to encode transaction")
	}

	return txBytes, nil
}

// isMuxedAccount returns true if the provided accountID is a muxed account (M...)
func isMuxedAccount(accountID string) bool {
	return strings.HasPrefix(accountID, "M")
}

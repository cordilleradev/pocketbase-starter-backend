package stellar

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"net/url"
// 	"strconv"
// 	"strings"

// 	"github.com/stellar/go/keypair"
// 	"github.com/stellar/go/strkey"
// 	"github.com/stellar/go/txnbuild"
// )

// // BuildSep10Transaction builds a SEP-10 challenge transaction.
// // The resulting transaction needs to be signed by the server and possibly also by the client,
// // depending on whether clientDomain is specified.
// //
// // Parameters:
// // - serverKeypair: The server's signing keypair (must be the signing key from the TOML file)
// // - clientAccountID: The stellar account ID that the client wishes to authenticate with
// // - homeDomain: The home domain of the service requiring authentication
// // - webAuthDomain: The domain of the service that requires authentication
// // - networkPassphrase: The network passphrase (test/public network)
// // - memo: Optional memo to attach (memos can be used for authentication of shared accounts)
// // - clientDomain: Optional domain of the client, used for client domain verification
// // - timebound: Optional custom timebound, defaults to 5 minute expiration
// //
// // Returns the SEP-10 challenge transaction in base64 XDR format, or an error if transaction creation fails.
// func BuildSep10Transaction(
// 	serverKeypair *keypair.Full,
// 	clientAccountID string,
// 	homeDomain string,
// 	webAuthDomain string,
// 	networkPassphrase string,
// 	memo string,
// 	clientDomain string,
// 	timebound *txnbuild.TimeBounds,
// ) (string, error) {
// 	// Validate client account ID
// 	clientAccountType, err := strkey.Version(clientAccountID)
// 	if err != nil {
// 		return "", fmt.Errorf("invalid client account ID: %w", err)
// 	}

// 	// If clientAccountID is muxed account but memo is present, return error
// 	if clientAccountType == strkey.VersionByteAccountID || clientAccountType == strkey.VersionByteAccount {
// 		if clientAccountType == strkey.VersionByteAccountID && memo != "" {
// 			return "", fmt.Errorf("memo cannot be used with muxed accounts")
// 		}
// 	} else {
// 		return "", fmt.Errorf("invalid account ID type: %v", clientAccountType)
// 	}

// 	// Validate homeDomain and create auth value
// 	if homeDomain == "" {
// 		return "", fmt.Errorf("home domain cannot be empty")
// 	}

// 	homedomain := homeDomain
// 	if !strings.HasSuffix(homedomain, " auth") {
// 		homedomain = homeDomain + " auth"
// 	}

// 	// Limit the home domain auth string to 64 characters
// 	if len(homedomain) > 64 {
// 		return "", fmt.Errorf("home domain string exceeds 64 characters: %s", homedomain)
// 	}

// 	// Check webAuthDomain
// 	if webAuthDomain == "" {
// 		return "", fmt.Errorf("web authentication domain cannot be empty")
// 	}

// 	if _, err := url.Parse("https://" + webAuthDomain); err != nil {
// 		return "", fmt.Errorf("invalid web authentication domain: %w", err)
// 	}

// 	// Generate a 64-byte value for the manageData operation
// 	randomBytes := make([]byte, 48)
// 	_, err = rand.Read(randomBytes)
// 	if err != nil {
// 		return "", fmt.Errorf("error generating random bytes: %w", err)
// 	}
// 	randomB64 := base64.StdEncoding.EncodeToString(randomBytes)

// 	// Create operations for the transaction
// 	operations := []txnbuild.Operation{
// 		&txnbuild.ManageData{
// 			SourceAccount: clientAccountID,
// 			Name:          homedomain,
// 			Value:         []byte(randomB64),
// 		},
// 		&txnbuild.ManageData{
// 			SourceAccount: serverKeypair.Address(),
// 			Name:          "web_auth_domain",
// 			Value:         []byte(webAuthDomain),
// 		},
// 	}

// 	// Setup timebound if not provided
// 	if timebound == nil {
// 		// 5 minute expiration is recommended in the SEP-10 spec
// 		timeout := txnbuild.NewTimeout(300) // 5 minutes in seconds
// 		timebound = &timeout
// 	}

// 	// Add client domain operation if needed
// 	if clientDomain != "" {
// 		// In a real implementation, you would fetch the SIGNING_KEY from the client domain's TOML file
// 		// For a proper implementation, replace this with code that fetches from the client domain
// 		// clientDomainKey = fetchClientDomainKey(clientDomain)

// 		// Add client domain operation
// 		operations = append(operations, &txnbuild.ManageData{
// 			SourceAccount: "", // This would be the client domain's SIGNING_KEY in a real implementation
// 			Name:          "client_domain",
// 			Value:         []byte(clientDomain),
// 		})
// 	}

// 	// Create transaction
// 	var txMemo txnbuild.Memo
// 	if memo != "" {
// 		// Only ID memos are supported in SEP-10
// 		memoID, err := strconv.ParseUint(memo, 10, 64)
// 		if err != nil {
// 			return "", fmt.Errorf("invalid memo ID format: %w", err)
// 		}
// 		txMemo = txnbuild.MemoID(memoID)
// 	} else {
// 		txMemo = txnbuild.MemoText("Proof of Ownership")
// 	}

// 	// Create transaction with sequence number 0 (this makes it invalid on the network)
// 	tx, err := txnbuild.NewTransaction(txnbuild.TransactionParams{
// 		SourceAccount: &txnbuild.SimpleAccount{
// 			AccountID: serverKeypair.Address(),
// 			Sequence:  0,
// 		},
// 		IncrementSequenceNum: false,
// 		Operations:           operations,
// 		BaseFee:              txnbuild.MinBaseFee,
// 		Memo:                 txMemo,
// 		Preconditions:        txnbuild.Preconditions{TimeBounds: *timebound},
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to build transaction: %w", err)
// 	}

// 	// Sign with the server signing key
// 	tx, err = tx.Sign(networkPassphrase, serverKeypair)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to sign transaction: %w", err)
// 	}

// 	// Get base64 encoded XDR string
// 	txeBase64, err := tx.Base64()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to encode transaction: %w", err)
// 	}

// 	return txeBase64, nil
// }

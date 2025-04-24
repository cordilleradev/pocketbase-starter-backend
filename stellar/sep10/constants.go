package sep10

// ErrSep10 represents an error in the SEP10 authentication flow
type ErrSep10 string

// Error method implements the error interface
func (e ErrSep10) Error() string {
	return string(e)
}

// SEP10 error codes
const (
	// Transaction creation errors
	ErrInvalidClientAccountID ErrSep10 = "invalid client account ID"
	ErrInvalidServerAccountID ErrSep10 = "invalid server account ID"
	ErrInvalidDomain          ErrSep10 = "invalid domain"
	ErrInvalidClientDomain    ErrSep10 = "invalid client domain"
	ErrInvalidHomeDomain      ErrSep10 = "invalid home domain"
	ErrInvalidWebAuthDomain   ErrSep10 = "invalid web auth domain"
	ErrMemoAndMuxedAccount    ErrSep10 = "cannot use memo with muxed account"
	ErrRandomGeneration       ErrSep10 = "could not generate random bytes"
	ErrHomeDomainTooLong      ErrSep10 = "home domain exceeds maximum length"
	ErrWebAuthDomainTooLong   ErrSep10 = "web auth domain exceeds maximum length"
	ErrClientDomainTooLong    ErrSep10 = "client domain exceeds maximum length"
	ErrInvalidClientDomainKP  ErrSep10 = "invalid client domain keypair"

	// Challenge validation errors
	ErrInvalidTransaction      ErrSep10 = "invalid challenge transaction"
	ErrInvalidSequenceNumber   ErrSep10 = "invalid sequence number"
	ErrInvalidTimeBounds       ErrSep10 = "invalid time bounds"
	ErrNoOperations            ErrSep10 = "no operations in challenge transaction"
	ErrInvalidOperationType    ErrSep10 = "invalid operation type in challenge transaction"
	ErrInvalidOperationSource  ErrSep10 = "invalid operation source account"
	ErrInvalidFirstOperation   ErrSep10 = "invalid first operation in challenge transaction"
	ErrInvalidSignature        ErrSep10 = "invalid signature in challenge transaction"
	ErrServerSignatureMissing  ErrSep10 = "server signature missing"
	ErrClientSignatureMissing  ErrSep10 = "client signature missing"
	ErrInsufficientWeight      ErrSep10 = "insufficient signature weight for required threshold"
	ErrClientAccountNotFound   ErrSep10 = "client account not found"
	ErrClientDomainNotVerified ErrSep10 = "client domain signature verification failed"
	ErrChallengeExpired        ErrSep10 = "challenge transaction has expired"
	ErrMultipleSignatures      ErrSep10 = "too many signatures provided"
)

// MaxHomeDomainLength is the maximum allowed home domain length in a SEP-10 challenge transaction
const MaxHomeDomainLength = 59 // 64 - len(" auth")

// MaxWebAuthDomainLength is the maximum allowed webauth domain length in a SEP-10 challenge transaction
const MaxWebAuthDomainLength = 64

// MaxClientDomainLength is the maximum allowed client domain length in a SEP-10 challenge transaction
const MaxClientDomainLength = 64

package sep1

// StellarToml represents the entire stellar.toml file structure
type StellarToml struct {
	Version                     string        `toml:"VERSION" json:"VERSION,omitempty"`
	NetworkPassphrase           string        `toml:"NETWORK_PASSPHRASE" json:"NETWORK_PASSPHRASE,omitempty"`
	FederationServer            string        `toml:"FEDERATION_SERVER" json:"FEDERATION_SERVER,omitempty"`
	AuthServer                  string        `toml:"AUTH_SERVER" json:"AUTH_SERVER,omitempty"`
	TransferServer              string        `toml:"TRANSFER_SERVER" json:"TRANSFER_SERVER,omitempty"`
	TransferServerSep0024       string        `toml:"TRANSFER_SERVER_SEP0024" json:"TRANSFER_SERVER_SEP0024,omitempty"`
	KYCServer                   string        `toml:"KYC_SERVER" json:"KYC_SERVER,omitempty"`
	WebAuthEndpoint             string        `toml:"WEB_AUTH_ENDPOINT" json:"WEB_AUTH_ENDPOINT,omitempty"`
	WebAuthForContractsEndpoint string        `toml:"WEB_AUTH_FOR_CONTRACTS_ENDPOINT" json:"WEB_AUTH_FOR_CONTRACTS_ENDPOINT,omitempty"`
	WebAuthContractID           string        `toml:"WEB_AUTH_CONTRACT_ID" json:"WEB_AUTH_CONTRACT_ID,omitempty"`
	SigningKey                  string        `toml:"SIGNING_KEY" json:"SIGNING_KEY,omitempty"`
	HorizonURL                  string        `toml:"HORIZON_URL" json:"HORIZON_URL,omitempty"`
	Accounts                    []string      `toml:"ACCOUNTS" json:"ACCOUNTS,omitempty"`
	URIRequestSigningKey        string        `toml:"URI_REQUEST_SIGNING_KEY" json:"URI_REQUEST_SIGNING_KEY,omitempty"`
	DirectPaymentServer         string        `toml:"DIRECT_PAYMENT_SERVER" json:"DIRECT_PAYMENT_SERVER,omitempty"`
	AnchorQuoteServer           string        `toml:"ANCHOR_QUOTE_SERVER" json:"ANCHOR_QUOTE_SERVER,omitempty"`
	Documentation               Documentation `toml:"DOCUMENTATION" json:"DOCUMENTATION"`
	Principals                  []Principal   `toml:"PRINCIPALS" json:"PRINCIPALS,omitempty"`
	Currencies                  []Currency    `toml:"CURRENCIES" json:"CURRENCIES,omitempty"`
	Validators                  []Validator   `toml:"VALIDATORS" json:"VALIDATORS,omitempty"`
}

// Documentation represents the organization documentation section
type Documentation struct {
	OrgName                       string `toml:"ORG_NAME" json:"ORG_NAME,omitempty"`
	OrgDBA                        string `toml:"ORG_DBA" json:"ORG_DBA,omitempty"`
	OrgURL                        string `toml:"ORG_URL" json:"ORG_URL,omitempty"`
	OrgLogo                       string `toml:"ORG_LOGO" json:"ORG_LOGO,omitempty"`
	OrgDescription                string `toml:"ORG_DESCRIPTION" json:"ORG_DESCRIPTION,omitempty"`
	OrgPhysicalAddress            string `toml:"ORG_PHYSICAL_ADDRESS" json:"ORG_PHYSICAL_ADDRESS,omitempty"`
	OrgPhysicalAddressAttestation string `toml:"ORG_PHYSICAL_ADDRESS_ATTESTATION" json:"ORG_PHYSICAL_ADDRESS_ATTESTATION,omitempty"`
	OrgPhoneNumber                string `toml:"ORG_PHONE_NUMBER" json:"ORG_PHONE_NUMBER,omitempty"`
	OrgPhoneNumberAttestation     string `toml:"ORG_PHONE_NUMBER_ATTESTATION" json:"ORG_PHONE_NUMBER_ATTESTATION,omitempty"`
	OrgKeybase                    string `toml:"ORG_KEYBASE" json:"ORG_KEYBASE,omitempty"`
	OrgTwitter                    string `toml:"ORG_TWITTER" json:"ORG_TWITTER,omitempty"`
	OrgGithub                     string `toml:"ORG_GITHUB" json:"ORG_GITHUB,omitempty"`
	OrgOfficialEmail              string `toml:"ORG_OFFICIAL_EMAIL" json:"ORG_OFFICIAL_EMAIL,omitempty"`
	OrgSupportEmail               string `toml:"ORG_SUPPORT_EMAIL" json:"ORG_SUPPORT_EMAIL,omitempty"`
	OrgLicensingAuthority         string `toml:"ORG_LICENSING_AUTHORITY" json:"ORG_LICENSING_AUTHORITY,omitempty"`
	OrgLicenseType                string `toml:"ORG_LICENSE_TYPE" json:"ORG_LICENSE_TYPE,omitempty"`
	OrgLicenseNumber              string `toml:"ORG_LICENSE_NUMBER" json:"ORG_LICENSE_NUMBER,omitempty"`
}

// Principal represents a principal point of contact
type Principal struct {
	Name                  string `toml:"name" json:"name,omitempty"`
	Email                 string `toml:"email" json:"email,omitempty"`
	Keybase               string `toml:"keybase" json:"keybase,omitempty"`
	Telegram              string `toml:"telegram" json:"telegram,omitempty"`
	Twitter               string `toml:"twitter" json:"twitter,omitempty"`
	Github                string `toml:"github" json:"github,omitempty"`
	IDPhotoHash           string `toml:"id_photo_hash" json:"id_photo_hash,omitempty"`
	VerificationPhotoHash string `toml:"verification_photo_hash" json:"verification_photo_hash,omitempty"`
}

// Currency represents a currency/token offered by the anchor
type Currency struct {
	Code                        string   `toml:"code" json:"code,omitempty"`
	Issuer                      string   `toml:"issuer" json:"issuer,omitempty"`
	Contract                    string   `toml:"contract" json:"contract,omitempty"`
	CodeTemplate                string   `toml:"code_template" json:"code_template,omitempty"`
	Status                      string   `toml:"status" json:"status,omitempty"`
	DisplayDecimals             int      `toml:"display_decimals" json:"display_decimals,omitempty"`
	Name                        string   `toml:"name" json:"name,omitempty"`
	Desc                        string   `toml:"desc" json:"desc,omitempty"`
	Conditions                  string   `toml:"conditions" json:"conditions,omitempty"`
	Image                       string   `toml:"image" json:"image,omitempty"`
	FixedNumber                 int      `toml:"fixed_number" json:"fixed_number,omitempty"`
	MaxNumber                   int      `toml:"max_number" json:"max_number,omitempty"`
	IsUnlimited                 bool     `toml:"is_unlimited" json:"is_unlimited,omitempty"`
	IsAssetAnchored             bool     `toml:"is_asset_anchored" json:"is_asset_anchored,omitempty"`
	AnchorAssetType             string   `toml:"anchor_asset_type" json:"anchor_asset_type,omitempty"`
	AnchorAsset                 string   `toml:"anchor_asset" json:"anchor_asset,omitempty"`
	AttestationOfReserve        string   `toml:"attestation_of_reserve" json:"attestation_of_reserve,omitempty"`
	RedemptionInstructions      string   `toml:"redemption_instructions" json:"redemption_instructions,omitempty"`
	CollateralAddresses         []string `toml:"collateral_addresses" json:"collateral_addresses,omitempty"`
	CollateralAddressMessages   []string `toml:"collateral_address_messages" json:"collateral_address_messages,omitempty"`
	CollateralAddressSignatures []string `toml:"collateral_address_signatures" json:"collateral_address_signatures,omitempty"`
	Regulated                   bool     `toml:"regulated" json:"regulated,omitempty"`
	ApprovalServer              string   `toml:"approval_server" json:"approval_server,omitempty"`
	ApprovalCriteria            string   `toml:"approval_criteria" json:"approval_criteria,omitempty"`
}

// Validator represents a validator node
type Validator struct {
	Alias       string `toml:"ALIAS" json:"ALIAS,omitempty"`
	DisplayName string `toml:"DISPLAY_NAME" json:"DISPLAY_NAME,omitempty"`
	PublicKey   string `toml:"PUBLIC_KEY" json:"PUBLIC_KEY,omitempty"`
	Host        string `toml:"HOST" json:"HOST,omitempty"`
	History     string `toml:"HISTORY" json:"HISTORY,omitempty"`
}

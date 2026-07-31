package client

type ValidatorLastBlock struct {
	Hash      string `json:"hash,omitempty"`
	Number    int64  `json:"number,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type WalletAttributes struct {
	Username                string              `json:"username,omitempty"`
	Vote                    string              `json:"vote,omitempty"`
	IsLegacy                bool                `json:"isLegacy,omitempty"`
	LegacyNonce             string              `json:"legacyNonce,omitempty"`
	ValidatorFee            BigInt              `json:"validatorFee,omitempty"`
	ValidatorRank           int64               `json:"validatorRank,omitempty"`
	ValidatorApproval       float64             `json:"validatorApproval,omitempty"`
	ValidatorResigned       bool                `json:"validatorResigned,omitempty"`
	ValidatorLastBlock      *ValidatorLastBlock `json:"validatorLastBlock,omitempty"`
	ValidatorPublicKey      string              `json:"validatorPublicKey,omitempty"`
	ValidatorForgedFees     BigInt              `json:"validatorForgedFees,omitempty"`
	ValidatorForgedTotal    BigInt              `json:"validatorForgedTotal,omitempty"`
	ValidatorVoteBalance    BigInt              `json:"validatorVoteBalance,omitempty"`
	ValidatorVotersCount    int64               `json:"validatorVotersCount,omitempty"`
	ValidatorForgedRewards  BigInt              `json:"validatorForgedRewards,omitempty"`
	ValidatorProducedBlocks int64               `json:"validatorProducedBlocks,omitempty"`
}

type Wallet struct {
	Address    string           `json:"address,omitempty"`
	PublicKey  string           `json:"publicKey,omitempty"`
	Balance    BigInt           `json:"balance,omitempty"`
	Nonce      BigInt           `json:"nonce,omitempty"`
	Attributes WalletAttributes `json:"attributes,omitempty"`
	UpdatedAt  string           `json:"updated_at,omitempty"`
	TokenCount int64            `json:"tokenCount,omitempty"`
}

type Wallets struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

type GetWallet struct {
	Data Wallet `json:"data,omitempty"`
}

package client

type LegacyColdWalletAttributes struct {
	LegacyNonce string `json:"legacyNonce,omitempty"`
}

type LegacyColdWallet struct {
	Address                  string                     `json:"address,omitempty"`
	PublicKey                string                     `json:"publicKey,omitempty"`
	Balance                  BigInt                     `json:"balance,omitempty"`
	Attributes               LegacyColdWalletAttributes `json:"attributes,omitempty"`
	MergeInfoWalletAddress   *string                    `json:"mergeInfoWalletAddress,omitempty"`
	MergeInfoTransactionHash *string                    `json:"mergeInfoTransactionHash,omitempty"`
}

type LegacyColdWallets struct {
	Meta Meta               `json:"meta,omitempty"`
	Data []LegacyColdWallet `json:"data,omitempty"`
}

type GetLegacyColdWallet struct {
	Data LegacyColdWallet `json:"data,omitempty"`
}

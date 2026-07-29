// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type ValidatorLastBlock struct {
	Id        string `json:"id,omitempty"`
	Height    int64  `json:"height,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type WalletAttributes struct {
	Username                string              `json:"username,omitempty"`
	Vote                    string              `json:"vote,omitempty"`
	ValidatorRank           byte                `json:"validatorRank,omitempty"`
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

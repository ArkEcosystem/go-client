// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Wallet struct {
	Address    string `json:"address,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
	Nonce      uint64 `json:"nonce,omitempty,string"`
	Balance    uint64 `json:"balance,omitempty,string"`
	IsDelegate bool   `json:"isDelegate,omitempty"`
}

type Wallets struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

type GetWallet struct {
	Data Wallet `json:"data,omitempty"`
}

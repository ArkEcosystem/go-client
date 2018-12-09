// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type BlockForged struct {
	Reward int64 `json:"reward,omitempty"`
	Fee    int64 `json:"fee,omitempty"`
	Total  int64 `json:"total,omitempty"`
}

type BlockPayload struct {
	Hash   string `json:"hash,omitempty"`
	Length byte   `json:"length,omitempty"`
}

type BlockGenerator struct {
	Username  string `json:"username,omitempty"`
	Address   string `json:"address,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type Block struct {
	Id           string         `json:"id,omitempty"`
	Version      byte           `json:"version,omitempty"`
	Height       int64          `json:"height,omitempty"`
	Previous     string         `json:"previous,omitempty"`
	Forged       BlockForged    `json:"forged,omitempty"`
	Payload      BlockPayload   `json:"payload,omitempty"`
	Generator    BlockGenerator `json:"generator,omitempty"`
	Signature    string         `json:"signature,omitempty"`
	Transactions byte           `json:"transactions,omitempty"`
	Timestamp    Timestamp      `json:"timestamp,omitempty"`
}

type Blocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

type GetBlock struct {
	Data Block `json:"data,omitempty"`
}

type GetBlockTransactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}

// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

type Block struct {
	Id       string `json:"id,omitempty"`
	Version  byte   `json:"version,omitempty"`
	Height   int64  `json:"height,omitempty"`
	Previous string `json:"previous,omitempty"`
	Forged   struct {
		Reward int64 `json:"reward,omitempty"`
		Fee    int64 `json:"fee,omitempty"`
		Total  int64 `json:"total,omitempty"`
	} `json:"forged,omitempty"`
	Payload struct {
		Hash   string `json:"hash,omitempty"`
		Lenght byte   `json:"length,omitempty"`
	} `json:"payload,omitempty"`
	Generator struct {
		Username  string `json:"username,omitempty"`
		Address   string `json:"address,omitempty"`
		PublicKey string `json:"publicKey,omitempty"`
	} `json:"generator,omitempty"`
	Signature    string    `json:"signature,omitempty"`
	Transactions byte      `json:"transactions,omitempty"`
	Timestamp    Timestamp `json:"timestamp,omitempty"`
}

type Blocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

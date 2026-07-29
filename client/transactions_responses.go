// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type TransactionReceipt struct {
	GasRefunded uint32 `json:"gasRefunded,omitempty"`
	GasUsed     uint32 `json:"gasUsed,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type Transaction struct {
	Hash            string             `json:"hash,omitempty"`
	Value           BigInt             `json:"value,omitempty"`
	BlockNumber     string             `json:"blockNumber,omitempty"`
	Confirmations   uint32             `json:"confirmations,omitempty"`
	Data            string             `json:"data,omitempty"`
	Gas             BigInt             `json:"gas,omitempty"`
	GasPrice        BigInt             `json:"gasPrice,omitempty"`
	Nonce           BigInt             `json:"nonce,omitempty"`
	To              string             `json:"to,omitempty"`
	From            string             `json:"from,omitempty"`
	SenderPublicKey string             `json:"senderPublicKey,omitempty"`
	Signature       string             `json:"signature,omitempty"`
	Timestamp       string             `json:"timestamp,omitempty"`
	Receipt         TransactionReceipt `json:"receipt,omitempty"`
}

type Transactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}

type GetTransaction struct {
	Data Transaction `json:"data,omitempty"`
}

type GetCreateTransaction struct {
	Data CreateTransaction `json:"data,omitempty"`
}

type TypeGroupTypes map[string]byte

type TransactionTypes struct {
	Data map[string]TypeGroupTypes `json:"data,omitempty"`
}

type TransactionFees struct {
	Data map[string]uint32 `json:"data,omitempty"`
}

type CreateTransaction struct {
	Accept  []string `json:"accept,omitempty"`
	Excess  []string `json:"excess,omitempty"`
	Invalid []string `json:"invalid,omitempty"`
}

// TransactionSchemas represents the response from the /transactions/schemas endpoint.
type TransactionSchemas struct {
	Data interface{} `json:"data"`
}

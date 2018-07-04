// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

type Transaction struct {
	Success         bool              `json:"success,omitempty"`
	Id              string            `json:"id,omitempty"`
	BlockId         string            `json:"blockid,omitempty"`
	Type            byte              `json:"type,omitempty"`
	Timestamp       int32             `json:"timestamp,omitempty"`
	Amount          int64             `json:"amount,omitempty"`
	Fee             int64             `json:"fee,omitempty"`
	RecipientId     string            `json:"recipientId,omitempty"`
	SenderId        string            `json:"senderId,omitempty"`
	SenderPublicKey string            `json:"senderPublicKey,omitempty"`
	VendorField     string            `json:"vendorField,omitempty"`
	Signature       string            `json:"signature,omitempty"`
	Asset           map[string]string `json:"asset,omitempty"`
	Confirmations   byte              `json:"confirmations,omitempty"`
}

type Transactions struct {
	Success      bool          `json:"success,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

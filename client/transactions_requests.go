// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type CreateTransactionRequest struct {
	Transactions []interface{} `json:"transactions,omitempty"`
}

type TransactionsSearchRequest struct {
	OrderBy         string  `json:"orderBy,omitempty"`
	Id              string  `json:"id,omitempty"`
	SenderId        string  `json:"senderId,omitempty"`
	BlockId         string  `json:"blockId,omitempty"`
	Type            uint8   `json:"type,omitempty"`
	Version         uint8   `json:"version,omitempty"`
	SenderPublicKey string  `json:"senderPublicKey,omitempty"`
	RecipientId     string  `json:"recipientId,omitempty"`
	Timestamp       *FromTo `json:"timestamp,omitempty"`
	Amount          *FromTo `json:"amount,omitempty"`
	Fee             *FromTo `json:"fee,omitempty"`
	VendorFieldHex  string  `json:"vendorFieldHex,omitempty"`
}

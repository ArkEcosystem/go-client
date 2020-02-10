// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Lock struct {
	LockId          string    `json:"lockId,omitempty"`
	Amount          uint64    `json:"amount,omitempty,string"`
	SecretHash      string    `json:"secretHash,omitempty"`
	SenderPublicKey string    `json:"senderPublicKey,omitempty"`
	RecipientId     string    `json:"recipientId,omitempty"`
	Timestamp       Timestamp `json:"timestamp,omitempty"`
	ExpirationType  byte      `json:"expirationType,omitempty"`
	ExpirationValue uint32    `json:"expirationValue,omitempty"`
	VendorField     string    `json:"vendorField,omitempty"`
	IsExpired       bool      `json:"isExpired,omitempty"`
}

type Locks struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Lock `json:"data,omitempty"`
}

type GetLock struct {
	Data Lock `json:"data,omitempty"`
}

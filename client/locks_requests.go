// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type LocksUnlockedRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type LocksSearchRequest struct {
	LockId          string  `json:"lockId,omitempty"`
	Amount          *FromTo `json:"amount,omitempty"`
	SecretHash      string  `json:"secretHash,omitempty"`
	SenderPublicKey string  `json:"senderPublicKey,omitempty"`
	RecipientId     string  `json:"recipientId,omitempty"`
	Timestamp       *FromTo `json:"timestamp,omitempty"`
	ExpirationType  byte    `json:"expirationType,omitempty"`
	ExpirationValue *FromTo `json:"expirationValue,omitempty"`
	VendorField     string  `json:"vendorField,omitempty"`
	IsExpired       bool    `json:"isExpired,omitempty"`
}

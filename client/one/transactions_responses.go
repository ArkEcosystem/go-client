// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

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

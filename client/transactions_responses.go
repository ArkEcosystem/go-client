// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"time"
)

type Transaction struct {
	Id              string            `json:"id,omitempty"`
	BlockId         string            `json:"blockId,omitempty"`
	Type            byte              `json:"type,omitempty"`
	TypeGroup       uint16            `json:"typeGroup,omitempty"`
	Amount          uint64            `json:"amount,omitempty,string"`
	Fee             uint64            `json:"fee,omitempty,string"`
	Sender          string            `json:"sender,omitempty"`
	SenderPublicKey string            `json:"senderPublicKey,omitempty"`
	Recipient       string            `json:"recipient,omitempty"`
	Signature       string            `json:"signature,omitempty"`
	VendorField     string            `json:"vendorField,omitempty"`
	Asset           *TransactionAsset `json:"asset,omitempty"`
	Confirmations   uint32            `json:"confirmations,omitempty"`
	Timestamp       Timestamp         `json:"timestamp,omitempty"`
	Nonce           string            `json:"nonce,omitempty"`
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

type Timestamp struct {
	Epoch int32  `json:"epoch,omitempty"`
	Unix  int32  `json:"unix,omitempty"`
	Human string `json:"human,omitempty"`
}

// Time parses the unix value of the timestamp and returns as time.Time object with
// location as local.
func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t.Unix), 0)
}

type CreateTransaction struct {
	Accept  []string `json:"accept,omitempty"`
	Excess  []string `json:"excess,omitempty"`
	Invalid []string `json:"invalid,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// TRANSACTION ASSETS //////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

type TransactionAsset struct {
	Votes          []string                          `json:"votes,omitempty"`
	Signature      *SecondSignatureRegistrationAsset `json:"signature,omitempty"`
	Delegate       *DelegateAsset                    `json:"publicKey,omitempty"`
	MultiSignature *MultiSignatureRegistrationAsset  `json:"multisignature,omitempty"`
	Ipfs           *IpfsAsset                        `json:"ipfs,omitempty"`
	Payments       []*MultiPaymentAsset              `json:"payments,omitempty"`
}

type SecondSignatureRegistrationAsset struct {
	PublicKey string `json:"publicKey,omitempty"`
}

type DelegateAsset struct {
	Username string `json:"username,omitempty"`
}

type MultiSignatureRegistrationAsset struct {
	Min       byte     `json:"min,omitempty"`
	Keysgroup []string `json:"keysgroup,omitempty"`
	Lifetime  byte     `json:"lifetime,omitempty"`
}

type IpfsAsset struct {
	Dag string `json:"dag,omitempty"`
}

type MultiPaymentAsset struct {
	Amount      uint64 `json:"amount,omitempty,string"`
	RecipientId string `json:"recipientId,omitempty"`
}

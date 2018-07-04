// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

type Transaction struct {
	Address       string            `json:"address,omitempty"`
	Id            string            `json:"id,omitempty"`
	BlockId       string            `json:"blockId,omitempty"`
	Type          byte              `json:"type,omitempty"`
	Amount        int64             `json:"amount,omitempty"`
	Fee           int64             `json:"fee,omitempty"`
	Sender        string            `json:"sender,omitempty"`
	Recipient     string            `json:"recipient,omitempty"`
	Signature     string            `json:"signature,omitempty"`
	VendorField   string            `json:"vendorField,omitempty"`
	Asset         map[string]string `json:"asset,omitempty"`
	Confirmations byte              `json:"confirmations,omitempty"`
	Timestamp     Timestamp         `json:"timestamp,omitempty"`
}

type Transactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}

type TransactionTypes struct {
	Data map[string]byte `json:"publicKey,omitempty"`
}

type Timestamp struct {
	Epoch int32  `json:"epoch,omitempty"`
	Unix  int32  `json:"unix,omitempty"`
	Human string `json:"human,omitempty"`
}

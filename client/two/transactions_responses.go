// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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

type CreateTransaction struct {
	Accept  map[string]string `json:"accept,omitempty"`
	Excess  map[string]string `json:"excess,omitempty"`
	Invalid map[string]string `json:"invalid,omitempty"`
}

// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type BlocksSearchRequest struct {
	Id                   string  `json:"id,omitempty"`
	Version              uint16  `json:"version,omitempty"`
	PreviousBlock        string  `json:"previousBlock,omitempty"`
	PayloadHash          string  `json:"payloadHash,omitempty"`
	GeneratorPublicKey   string  `json:"generatorPublicKey,omitempty"`
	BlockSignature       string  `json:"blockSignature,omitempty"`
	Timestamp            *FromTo `json:"timestamp,omitempty"`
	Height               *FromTo `json:"height,omitempty"`
	NumberOfTransactions *FromTo `json:"numberOfTransactions,omitempty"`
	TotalAmount          *FromTo `json:"totalAmount,omitempty"`
	TotalFee             *FromTo `json:"totalFee,omitempty"`
	Reward               *FromTo `json:"reward,omitempty"`
	PayloadLength        *FromTo `json:"payloadLength,omitempty"`
}

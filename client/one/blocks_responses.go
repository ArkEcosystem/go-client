// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type Block struct {
	Success              bool   `json:"success,omitempty"`
	Id                   string `json:"id,omitempty"`
	Version              int    `json:"version,omitempty"`
	Timestamp            int32  `json:"timestamp,omitempty"`
	PreviousBlock        string `json:"previousBlock,omitempty"`
	Height               int    `json:"height,omitempty"`
	NumberOfTransactions int    `json:"numberOfTransactions,omitempty"`
	TotalAmount          int64  `json:"totalAmount,omitempty"`
	TotalForged          string `json:"totalForged,omitempty"`
	TotalFee             int64  `json:"totalFee,omitempty"`
	Reward               int64  `json:"reward,omitempty"`
	PayloadLength        int    `json:"payloadLength,omitempty"`
	PayloadHash          string `json:"payloadHash,omitempty"`
	GeneratorPublicKey   string `json:"generatorPublicKey,omitempty"`
	BlockSignature       string `json:"blockSignature,omitempty"`
	Confirmations        int    `json:"confirmations,omitempty"`
}

type Blocks struct {
	Success bool    `json:"success,omitempty"`
	Blocks  []Block `json:"blocks,omitempty"`
}

type BlocksEpoch struct {
	Success bool `json:"success,omitempty"`
	Epoch   int  `json:"epoch,omitempty"`
}

type BlocksFee struct {
	Success bool  `json:"success,omitempty"`
	Fee     int64 `json:"fee,omitempty"`
}

type BlocksFees struct {
	Success bool `json:"success,omitempty"`
	Fees    struct {
		Send            int64 `json:"send,omitempty"`
		Vote            int64 `json:"vote,omitempty"`
		Secondsignature int64 `json:"secondsignature,omitempty"`
		Delegate        int64 `json:"delegate,omitempty"`
		Multisignature  int64 `json:"multisignature,omitempty"`
	} `json:"fees,omitempty"`
}

type BlocksHeight struct {
	Success bool  `json:"success,omitempty"`
	Height  int64 `json:"height,omitempty"`
}

type BlocksMilestone struct {
	Success   bool `json:"success,omitempty"`
	Milestone int  `json:"milestone,omitempty"`
}

type BlocksNethash struct {
	Success bool `json:"success,omitempty"`
	Nethash int  `json:"nethash,omitempty"`
}

type BlocksReward struct {
	Success bool  `json:"success,omitempty"`
	Reward  int64 `json:"reward,omitempty"`
}

type BlocksStatus struct {
	Success   bool   `json:"success,omitempty"`
	Epoch     string `json:"epoch,omitempty"`
	Height    int64  `json:"height,omitempty"`
	Fee       int64  `json:"fee,omitempty"`
	Milestone int    `json:"milestone,omitempty"`
	Nethash   string `json:"nethash,omitempty"`
	Reward    int64  `json:"reward,omitempty"`
	Supply    int64  `json:"supply,omitempty"`
}

type BlocksSupply struct {
	Success bool  `json:"success,omitempty"`
	Supply  int64 `json:"count,omitempty"`
}

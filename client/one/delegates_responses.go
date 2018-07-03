// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type Delegate struct {
	Success        bool    `json:"success,omitempty"`
	Username       string  `json:"username,omitempty"`
	Address        string  `json:"address,omitempty"`
	PublicKey      string  `json:"publicKey,omitempty"`
	Vote           string  `json:"vote,omitempty"`
	ProducedBlocks int     `json:"producedblocks,omitempty"`
	MissedBlocks   int     `json:"missedblocks,omitempty"`
	Rate           int     `json:"rate,omitempty"`
	Approval       float64 `json:"approval,omitempty"`
	Productivity   float64 `json:"productivity,omitempty"`
}

type DelegatesCount struct {
	Success bool `json:"success,omitempty"`
	Count   int  `json:"count,omitempty"`
}

type DelegateFee struct {
	Success bool  `json:"success,omitempty"`
	Fee     int64 `json:"fee,omitempty"`
}

type NextForger struct {
	Success      bool              `json:"success,omitempty"`
	CurrentBlock int               `json:"currentBlock,omitempty"`
	CurrentSlot  int64             `json:"currentSlot,omitempty"`
	Delegates    map[string]string `json:"delegates,omitempty"`
}

type ForgedByDelegate struct {
	Success bool  `json:"success,omitempty"`
	Fees    int64 `json:"fees,omitempty"`
	Rewards int64 `json:"rewards,omitempty"`
	Forged  int64 `json:"forged,omitempty"`
}

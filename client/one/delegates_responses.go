// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

type Delegate struct {
	Success        bool    `json:"success,omitempty"`
	Username       string  `json:"username,omitempty"`
	Address        string  `json:"address,omitempty"`
	PublicKey      string  `json:"publicKey,omitempty"`
	Vote           string  `json:"vote,omitempty"`
	ProducedBlocks byte    `json:"producedblocks,omitempty"`
	MissedBlocks   byte    `json:"missedblocks,omitempty"`
	Rate           byte    `json:"rate,omitempty"`
	Approval       float64 `json:"approval,omitempty"`
	Productivity   float64 `json:"productivity,omitempty"`
}

type GetDelegate struct {
	Success  bool     `json:"success,omitempty"`
	Delegate Delegate `json:"delegate,omitempty"`
}

type DelegatesCount struct {
	Success bool `json:"success,omitempty"`
	Count   byte `json:"count,omitempty"`
}

type DelegateFee struct {
	Success bool  `json:"success,omitempty"`
	Fee     int64 `json:"fee,omitempty"`
}

type NextForger struct {
	Success      bool     `json:"success,omitempty"`
	CurrentBlock byte     `json:"currentBlock,omitempty"`
	CurrentSlot  int64    `json:"currentSlot,omitempty"`
	Delegates    []string `json:"delegates,omitempty"`
}

type ForgedByDelegate struct {
	Success bool   `json:"success,omitempty"`
	Fees    string `json:"fees,omitempty"`
	Rewards string `json:"rewards,omitempty"`
	Forged  string `json:"forged,omitempty"`
}

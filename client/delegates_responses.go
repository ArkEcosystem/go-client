// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type DelegateBlocks struct {
	Produced uint32 `json:"produced,omitempty"`
	Missed   uint32 `json:"missed,omitempty"`
	Last     Block  `json:"last,omitempty"`
}

type DelegateProduction struct {
	Approval float64 `json:"approval,omitempty"`
}

type DelegateForged struct {
	Fees    uint64 `json:"fees,omitempty,string"`
	Rewards uint64 `json:"rewards,omitempty,string"`
	Total   uint64 `json:"total,omitempty,string"`
}

type Delegate struct {
	Username   string             `json:"username,omitempty"`
	Address    string             `json:"address,omitempty"`
	PublicKey  string             `json:"publicKey,omitempty"`
	Votes      int64              `json:"votes,omitempty"`
	Rank       byte               `json:"rank,omitempty"`
	Blocks     DelegateBlocks     `json:"blocks,omitempty"`
	Production DelegateProduction `json:"production,omitempty"`
	Forged     DelegateForged     `json:"forged,omitempty"`
}

type Delegates struct {
	Meta Meta       `json:"meta,omitempty"`
	Data []Delegate `json:"data,omitempty"`
}

type GetDelegate struct {
	Meta Meta     `json:"meta,omitempty"`
	Data Delegate `json:"data,omitempty"`
}

type GetDelegateBlocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

type GetDelegateVoters struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

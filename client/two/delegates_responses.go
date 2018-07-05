// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

type DelegateBlocks struct {
	Produced string `json:"produced,omitempty"`
	Missed   string `json:"missed,omitempty"`
	Last     Block  `json:"last,omitempty"`
}

type DelegateProduction struct {
	Approval     string `json:"approval,omitempty"`
	Productivity string `json:"productivity,omitempty"`
}

type Delegate struct {
	Username   string             `json:"username,omitempty"`
	Address    string             `json:"address,omitempty"`
	PublicKey  string             `json:"publicKey,omitempty"`
	Votes      int64              `json:"votes,omitempty"`
	Rank       byte               `json:"rank,omitempty"`
	Blocks     DelegateBlocks     `json:"blocks,omitempty"`
	Production DelegateProduction `json:"production,omitempty"`
}

type Delegates struct {
	Meta Meta       `json:"meta,omitempty"`
	Data []Delegate `json:"data,omitempty"`
}

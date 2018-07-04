// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

type Delegate struct {
	Username  string `json:"username,omitempty"`
	Address   string `json:"address,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Votes     int64  `json:"votes,omitempty"`
	Rank      byte   `json:"rank,omitempty"`
	Blocks    struct {
		Produced string `json:"produced,omitempty"`
		Missed   string `json:"missed,omitempty"`
		Last     Block  `json:"last,omitempty"`
	} `json:"blocks,omitempty"`
	Production struct {
		Approval     string `json:"approval,omitempty"`
		Productivity string `json:"productivity,omitempty"`
	} `json:"production,omitempty"`
}

type Delegates struct {
	Meta Meta       `json:"meta,omitempty"`
	Data []Delegate `json:"data,omitempty"`
}

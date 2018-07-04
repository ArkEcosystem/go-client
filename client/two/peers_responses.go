// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

type Peer struct {
	Address string `json:"address,omitempty"`
	ip      string `json:"ip,omitempty"`
	port    byte   `json:"port,omitempty"`
	version string `json:"version,omitempty"`
	height  byte   `json:"height,omitempty"`
	status  string `json:"status,omitempty"`
	os      string `json:"os,omitempty"`
	latency byte   `json:"latency,omitempty"`
}

type Peers struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Peer `json:"data,omitempty"`
}

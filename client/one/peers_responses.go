// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type Peer struct {
	Success bool   `json:"success,omitempty"`
	Ip      string `json:"ip,omitempty"`
	Port    byte   `json:"port,omitempty"`
	Version string `json:"version,omitempty"`
	Errors  byte   `json:"errors,omitempty"`
	Os      string `json:"os,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Status  string `json:"status,omitempty"`
	Delay   byte   `json:"delay,omitempty"`
}

type Peers struct {
	Success bool   `json:"success,omitempty"`
	Peers   []Peer `json:"peers,omitempty"`
}

type PeersVersion struct {
	Success bool   `json:"success,omitempty"`
	Version string `json:"version,omitempty"`
}

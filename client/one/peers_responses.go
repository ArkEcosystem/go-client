// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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

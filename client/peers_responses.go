// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type PeerPorts map[string]int16

type Peer struct {
	Ip      string    `json:"ip,omitempty"`
	Port    int16     `json:"port,omitempty"`
	Ports   PeerPorts `json:"ports,omitempty"`
	Version string    `json:"version,omitempty"`
	Height  int64     `json:"height,omitempty"`
	Latency byte      `json:"latency,omitempty"`
}

type Peers struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Peer `json:"data,omitempty"`
}

type GetPeer struct {
	Meta Meta `json:"meta,omitempty"`
	Data Peer `json:"data,omitempty"`
}

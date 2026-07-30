// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type PeerPorts map[string]int64

type PeerPlugin struct {
	Enabled            bool  `json:"enabled,omitempty"`
	EstimateTotalCount bool  `json:"estimateTotalCount,omitempty"`
	Port               int64 `json:"port,omitempty"`
}

type Peer struct {
	BlockNumber int64                 `json:"blockNumber,omitempty"`
	Ip          string                `json:"ip,omitempty"`
	Latency     int64                 `json:"latency,omitempty"`
	Plugins     map[string]PeerPlugin `json:"plugins,omitempty"`
	Port        int64                 `json:"port,omitempty"`
	Ports       PeerPorts             `json:"ports,omitempty"`
	Version     string                `json:"version,omitempty"`
}

type Peers struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Peer `json:"data,omitempty"`
}

type GetPeer struct {
	Meta Meta `json:"meta,omitempty"`
	Data Peer `json:"data,omitempty"`
}

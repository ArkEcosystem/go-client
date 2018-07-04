// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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

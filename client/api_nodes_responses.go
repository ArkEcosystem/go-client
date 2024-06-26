// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type ApiNodesResponse struct {
	Meta Meta       `json:"meta"`
	Data []ApiNode  `json:"data"`
}

type ApiNode struct {
	Ip      string                 `json:"ip"`
	Port    int                    `json:"port"`
	Version string                 `json:"version,omitempty"`
	Height  int                    `json:"height,omitempty"`
	Latency int                    `json:"latency,omitempty"`
	Ports   map[string]interface{} `json:"ports,omitempty"`
	Plugins map[string]interface{} `json:"plugins,omitempty"`
}

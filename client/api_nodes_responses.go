// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type ApiNode struct {
	Url     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Latency int64  `json:"latency,omitempty"`
	Status  string `json:"status,omitempty"`
}

type ApiNodesResponse struct {
	Meta Meta      `json:"meta"`
	Data []ApiNode `json:"data"`
}

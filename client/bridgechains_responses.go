// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Bridgechain struct {
	BridgechainId         uint16   `json:"bridgechainId,omitempty"`
	BusinessId            uint16   `json:"businessId,omitempty"`
	Name                  string   `json:"name,omitempty"`
	SeedNodes             []string `json:"seedNodes,omitempty"`
	GenesisHash           string   `json:"genesisHash,omitempty"`
	BridgechainRepository string   `json:"bridgechainRepository,omitempty"`
}

type Bridgechains struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Bridgechain `json:"data,omitempty"`
}

type GetBridgechain struct {
	Data Bridgechain `json:"data,omitempty"`
}

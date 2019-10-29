// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type BridgechainsSearchRequest struct {
	BridgechainId         uint16   `json:"bridgechainId,omitempty"`
	BusinessId            uint16   `json:"businessId,omitempty"`
	Name                  string   `json:"name,omitempty"`
	SeedNodes             []string `json:"website,omitempty"`
	GenesisHash           string   `json:"vat,omitempty"`
	BridgechainRepository string   `json:"repository,omitempty"`
}

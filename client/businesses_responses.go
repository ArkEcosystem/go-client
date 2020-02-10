// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Business struct {
	BusinessId uint16 `json:"businessId,omitempty"`
	Name       string `json:"name,omitempty"`
	Website    string `json:"website,omitempty"`
	Vat        string `json:"vat,omitempty"`
	Repository string `json:"repository,omitempty"`
}

type Businesses struct {
	Meta Meta       `json:"meta,omitempty"`
	Data []Business `json:"data,omitempty"`
}

type GetBusiness struct {
	Data Business `json:"data,omitempty"`
}

type GetBusinessBridgechains struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Bridgechain `json:"data,omitempty"`
}

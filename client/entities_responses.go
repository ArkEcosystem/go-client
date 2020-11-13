// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Entity struct {
	Id         string    `json:"id,omitempty"`
	Address    string    `json:"address,omitempty"`
	PublicKey  string    `json:"publicKey,omitempty"`
	IsResigned bool      `json:"isResigned,omitempty"`
	Type       uint16    `json:"type,omitempty"`
	SubType    uint16    `json:"subType,omitempty"`
	Data       []Data    `json:"data,omitempty"`
}

type Entities struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Entity `json:"data,omitempty"`
}

type GetEntity struct {
	Data Entity `json:"data,omitempty"`
}

type Data struct {
	Name     string `json:"name,omitempty"`
	IpfsData string `json:"ipfsData,omitempty"`
}
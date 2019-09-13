// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type RoundDelegate struct {
	PublicKey string `json:"publicKey,omitempty"`
	Votes     string `json:"votes,omitempty"`
}

type GetDelegates struct {
	Data []RoundDelegate `json:"data,omitempty"`
}

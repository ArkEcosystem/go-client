// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

// CommitResponse represents the response from the /api/commits/<height> endpoint.
type CommitResponse struct {
	Data CommitData `json:"data"`
}

// CommitData represents the data field in the commit response.
type CommitData struct {
	Height     string   `json:"height"`
	Signature  string   `json:"signature"`
	Validators []string `json:"validators"`
}

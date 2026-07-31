// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

// BlockchainInfo represents the response from the /blockchain endpoint.
type BlockchainInfo struct {
	Data BlockchainData `json:"data"`
}

// BlockchainData represents the data field in the BlockchainInfo response.
type BlockchainData struct {
	Block  BlockchainBlock `json:"block"`
	Supply BigInt          `json:"supply"`
}

// BlockchainBlock represents a block in the blockchain.
type BlockchainBlock struct {
	Hash   string `json:"hash"`
	Number int64  `json:"number"`
}

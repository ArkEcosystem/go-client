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

package client

type VotesQuery struct {
	Pagination
	Address          string `url:"address,omitempty"`
	BlockHash        string `url:"blockHash,omitempty"`
	Data             string `url:"data,omitempty"`
	From             string `url:"from,omitempty"`
	GasPrice         string `url:"gasPrice,omitempty"`
	Hash             string `url:"hash,omitempty"`
	Nonce            int64  `url:"nonce,omitempty"`
	SenderPublicKey  string `url:"senderPublicKey,omitempty"`
	Timestamp        int64  `url:"timestamp,omitempty"`
	To               string `url:"to,omitempty"`
	TransactionIndex int64  `url:"transactionIndex,omitempty"`
	Value            string `url:"value,omitempty"`
	FullReceipt      bool   `url:"fullReceipt,omitempty"`
	OrderBy          string `url:"orderBy,omitempty"`
}

type VoteGetQuery struct {
	FullReceipt bool `url:"fullReceipt,omitempty"`
}

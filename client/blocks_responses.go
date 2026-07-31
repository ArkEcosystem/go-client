package client

type Block struct {
	Hash              string `json:"hash,omitempty"`
	Number            int64  `json:"number,omitempty"`
	Round             int64  `json:"round,omitempty"`
	Confirmations     uint32 `json:"confirmations,omitempty"`
	Amount            BigInt `json:"amount,omitempty"`
	Fee               BigInt `json:"fee,omitempty"`
	GasUsed           int64  `json:"gasUsed,omitempty"`
	Reward            BigInt `json:"reward,omitempty"`
	Total             BigInt `json:"total,omitempty"`
	Proposer          string `json:"proposer,omitempty"`
	PublicKey         string `json:"publicKey,omitempty"`
	Username          string `json:"username,omitempty"`
	ValidatorSet      string `json:"validatorSet,omitempty"`
	TransactionsRoot  string `json:"transactionsRoot,omitempty"`
	PayloadSize       int64  `json:"payloadSize,omitempty"`
	ParentHash        string `json:"parentHash,omitempty"`
	Signature         string `json:"signature,omitempty"`
	Timestamp         string `json:"timestamp,omitempty"`
	TransactionsCount int64  `json:"transactionsCount,omitempty"`
	Version           byte   `json:"version,omitempty"`
}

type Blocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

type GetBlock struct {
	Data Block `json:"data,omitempty"`
}

type GetBlockTransactions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []Transaction `json:"data,omitempty"`
}

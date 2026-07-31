package client

type ReceiptLog struct {
	Data    string   `json:"data,omitempty"`
	Topics  []string `json:"topics,omitempty"`
	Address string   `json:"address,omitempty"`
}

type Receipt struct {
	TransactionHash string       `json:"transactionHash,omitempty"`
	Status          int          `json:"status,omitempty"`
	GasUsed         uint32       `json:"gasUsed,omitempty"`
	GasRefunded     uint32       `json:"gasRefunded,omitempty"`
	ContractAddress *string      `json:"contractAddress,omitempty"`
	Logs            []ReceiptLog `json:"logs,omitempty"`
	Output          string       `json:"output,omitempty"`
}

type Receipts struct {
	Meta Meta      `json:"meta,omitempty"`
	Data []Receipt `json:"data,omitempty"`
}

type GetReceipt struct {
	Data Receipt `json:"data,omitempty"`
}

package client

type ReceiptsQuery struct {
	Pagination
	From            string `url:"from,omitempty"`
	To              string `url:"to,omitempty"`
	FullReceipt     bool   `url:"fullReceipt,omitempty"`
	IncludeTokens   bool   `url:"includeTokens,omitempty"`
	TransactionHash string `url:"transactionHash,omitempty"`
}

type ReceiptQuery struct {
	FullReceipt   bool `url:"fullReceipt,omitempty"`
	IncludeTokens bool `url:"includeTokens,omitempty"`
}

type ReceiptContractsQuery struct {
	Pagination
	From          string `url:"from,omitempty"`
	FullReceipt   bool   `url:"fullReceipt,omitempty"`
	IncludeTokens bool   `url:"includeTokens,omitempty"`
}

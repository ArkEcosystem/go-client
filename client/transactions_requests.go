// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type CreateTransactionRequest struct {
	Transactions []string `json:"transactions,omitempty"`
}

type TransactionsQuery struct {
	Pagination
	Address          string `url:"address,omitempty"`
	BlockHash        string `url:"blockHash,omitempty"`
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
	IncludeTokens    bool   `url:"includeTokens,omitempty"`
	OrderBy          string `url:"orderBy,omitempty"`
}

type UnconfirmedTransactionsQuery struct {
	Pagination
	OrderBy string `url:"orderBy,omitempty"`
}

type TransactionGetQuery struct {
	FullReceipt   bool `url:"fullReceipt,omitempty"`
	IncludeTokens bool `url:"includeTokens,omitempty"`
}

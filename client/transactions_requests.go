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

// TransactionsQuery is the query for Transactions.List and is also reused by
// Wallets.Transactions/SentTransactions/ReceivedTransactions/Votes.
// typescript-client also declares an "asset" field here, but live-server
// testing showed it's typed wrong there (declared as a string, but the API
// requires an object/array and 422s on any string value) - omitted here
// rather than shipping a field that can never work.
type TransactionsQuery struct {
	Pagination
	Address          string `url:"address,omitempty"`
	BlockHash        string `url:"blockHash,omitempty"`
	From             string `url:"from,omitempty"`
	GasPrice         string `url:"gasPrice,omitempty"`
	Hash             string `url:"hash,omitempty"`
	Nonce            int64  `url:"nonce,omitempty"`
	SenderId         string `url:"senderId,omitempty"`
	SenderPublicKey  string `url:"senderPublicKey,omitempty"`
	Timestamp        int64  `url:"timestamp,omitempty"`
	To               string `url:"to,omitempty"`
	TransactionIndex int64  `url:"transactionIndex,omitempty"`
	Value            string `url:"value,omitempty"`
	FullReceipt      bool   `url:"fullReceipt,omitempty"`
	IncludeTokens    bool   `url:"includeTokens,omitempty"`
	OrderBy          string `url:"orderBy,omitempty"`
}

// UnconfirmedTransactionsQuery is the query for Transactions.ListUnconfirmed.
type UnconfirmedTransactionsQuery struct {
	Pagination
	OrderBy string `url:"orderBy,omitempty"`
}

// TransactionGetQuery is the query for Transactions.Get.
type TransactionGetQuery struct {
	FullReceipt   bool `url:"fullReceipt,omitempty"`
	IncludeTokens bool `url:"includeTokens,omitempty"`
}

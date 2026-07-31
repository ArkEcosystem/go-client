package client

type TokensQuery struct {
	Pagination
	IgnoreWhitelist bool           `url:"ignoreWhitelist,omitempty"`
	Name            string         `url:"name,omitempty"`
	Whitelist       CommaSeparated `url:"whitelist,omitempty"`
}

type TokenLookupQuery struct {
	Pagination
	From            string `url:"from,omitempty"`
	To              string `url:"to,omitempty"`
	TransactionHash string `url:"transactionHash,omitempty"`
}

type TokenTransfersQuery struct {
	TokenLookupQuery
	Addresses       CommaSeparated `url:"addresses,omitempty"`
	IgnoreWhitelist bool           `url:"ignoreWhitelist,omitempty"`
	Whitelist       CommaSeparated `url:"whitelist,omitempty"`
}

type TokenApprovalsQuery struct {
	TokenLookupQuery
	Addresses CommaSeparated `url:"addresses,omitempty"`
	Whitelist CommaSeparated `url:"whitelist,omitempty"`
}

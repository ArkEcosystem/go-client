package client

type WalletTokensQuery struct {
	Addresses CommaSeparated `url:"addresses,omitempty"`
	Page      int            `url:"page,omitempty"`
	Limit     int            `url:"limit,omitempty"`
}

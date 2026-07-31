package client

type WalletsQuery struct {
	Pagination
	Address     string `url:"address,omitempty"`
	PublicKey   string `url:"publicKey,omitempty"`
	Balance     string `url:"balance,omitempty"`
	BalanceFrom string `url:"balance.from,omitempty"`
	BalanceTo   string `url:"balance.to,omitempty"`
	Nonce       int64  `url:"nonce,omitempty"`
	NonceFrom   int64  `url:"nonce.from,omitempty"`
	NonceTo     int64  `url:"nonce.to,omitempty"`
	Attributes  string `url:"attributes,omitempty"`
	OrderBy     string `url:"orderBy,omitempty"`
}

type WalletTokensForQuery struct {
	Pagination
	IgnoreWhitelist bool           `url:"ignoreWhitelist,omitempty"`
	MinBalance      float64        `url:"minBalance,omitempty"`
	Name            string         `url:"name,omitempty"`
	Whitelist       CommaSeparated `url:"whitelist,omitempty"`
}

type WalletTokensQuery struct {
	Pagination
	Addresses       CommaSeparated `url:"addresses,omitempty"`
	IgnoreWhitelist bool           `url:"ignoreWhitelist,omitempty"`
	MinBalance      float64        `url:"minBalance,omitempty"`
	Whitelist       CommaSeparated `url:"whitelist,omitempty"`
}

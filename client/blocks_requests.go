package client

type BlocksQuery struct {
	Pagination
	OrderBy       string `url:"orderBy,omitempty"`
	Hash          string `url:"hash,omitempty"`
	Number        int64  `url:"number,omitempty"`
	NumberFrom    int64  `url:"number.from,omitempty"`
	NumberTo      int64  `url:"number.to,omitempty"`
	Timestamp     int64  `url:"timestamp,omitempty"`
	TimestampFrom int64  `url:"timestamp.from,omitempty"`
	TimestampTo   int64  `url:"timestamp.to,omitempty"`
}

type BlockTransactionsQuery struct {
	Pagination
	OrderBy string `url:"orderBy,omitempty"`
}

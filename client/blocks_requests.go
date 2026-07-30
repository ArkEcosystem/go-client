package client

type BlocksQuery struct {
	Pagination
	OrderBy       string `url:"orderBy,omitempty"`
	Id            string `url:"id,omitempty"`
	Number        int64  `url:"number,omitempty"`
	Height        int64  `url:"height,omitempty"`
	HeightFrom    int64  `url:"height.from,omitempty"`
	HeightTo      int64  `url:"height.to,omitempty"`
	Timestamp     int64  `url:"timestamp,omitempty"`
	TimestampFrom int64  `url:"timestamp.from,omitempty"`
	TimestampTo   int64  `url:"timestamp.to,omitempty"`
}

type BlockTransactionsQuery struct {
	Pagination
	OrderBy string `url:"orderBy,omitempty"`
}

package client

type ApiNodesQuery struct {
	Pagination
	Ip      string `url:"ip,omitempty"`
	Version string `url:"version,omitempty"`
	OrderBy string `url:"orderBy,omitempty"`
}

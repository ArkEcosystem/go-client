package client

type PeersQuery struct {
	Pagination
	Ip      string `url:"ip,omitempty"`
	OrderBy string `url:"orderBy,omitempty"`
	Version string `url:"version,omitempty"`
}

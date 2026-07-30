package client

// ApiNodesQuery is the query for ApiNodes.All.
type ApiNodesQuery struct {
	Pagination
	Ip      string `url:"ip,omitempty"`
	Version string `url:"version,omitempty"`
	OrderBy string `url:"orderBy,omitempty"`
}

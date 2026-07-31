package client

type ApiNode struct {
	Url     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Latency int64  `json:"latency,omitempty"`
	Status  string `json:"status,omitempty"`
}

type ApiNodesResponse struct {
	Meta Meta      `json:"meta"`
	Data []ApiNode `json:"data"`
}

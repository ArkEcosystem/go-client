package client

// CommitResponse represents the response from the /api/commits/<height> endpoint.
type CommitResponse struct {
	Data CommitData `json:"data"`
}

// CommitData represents the data field in the commit response.
type CommitData struct {
	BlockNumber string   `json:"blockNumber"`
	Signature   string   `json:"signature"`
	Validators  []string `json:"validators"`
}

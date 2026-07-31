package client

// RoundData represents the data field in the round response.
type RoundData struct {
	Round       string   `json:"round"`
	RoundHeight string   `json:"roundHeight"`
	Validators  []string `json:"validators"`
	Votes       []string `json:"votes"`
}

// GetRounds represents the response from the /api/rounds endpoint.
type GetRounds struct {
	Meta Meta        `json:"meta"`
	Data []RoundData `json:"data"`
}

// GetRound represents the response from the /api/rounds/<id> endpoint.
type GetRound struct {
	Data RoundData `json:"data"`
}

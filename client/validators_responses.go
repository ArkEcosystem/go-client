package client

type ValidatorBlocks struct {
	Produced uint32 `json:"produced,omitempty"`
	Missed   uint32 `json:"missed,omitempty"`
	Last     Block  `json:"last,omitempty"`
}

type ValidatorProduction struct {
	Approval float64 `json:"approval,omitempty"`
}

type ValidatorForged struct {
	Fees    uint64 `json:"fees,omitempty,string"`
	Rewards uint64 `json:"rewards,omitempty,string"`
	Total   uint64 `json:"total,omitempty,string"`
}

type Validator struct {
	Username   string              `json:"username,omitempty"`
	Address    string              `json:"address,omitempty"`
	PublicKey  string              `json:"publicKey,omitempty"`
	Votes      int64               `json:"votes,omitempty"`
	Rank       byte                `json:"rank,omitempty"`
	Blocks     ValidatorBlocks     `json:"blocks,omitempty"`
	Production ValidatorProduction `json:"production,omitempty"`
	Forged     ValidatorForged     `json:"forged,omitempty"`
}

type Validators struct {
	Meta Meta        `json:"meta,omitempty"`
	Data []Validator `json:"data,omitempty"`
}

type GetValidator struct {
	Meta Meta      `json:"meta,omitempty"`
	Data Validator `json:"data,omitempty"`
}

type GetValidatorBlocks struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Block `json:"data,omitempty"`
}

type GetValidatorVoters struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

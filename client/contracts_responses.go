package client

type Contract struct {
	ActiveImplementation string   `json:"activeImplementation,omitempty"`
	Address              string   `json:"address,omitempty"`
	Implementations      []string `json:"implementations,omitempty"`
	Proxy                string   `json:"proxy,omitempty"`
}

type ContractsResponse struct {
	Data map[string]Contract `json:"data,omitempty"`
}

type ContractAbi struct {
	Abi []interface{} `json:"abi,omitempty"`
}

type ContractAbiResponse struct {
	Data ContractAbi `json:"data,omitempty"`
}

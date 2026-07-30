package client

type Token struct {
	Address        string `json:"address,omitempty"`
	Symbol         string `json:"symbol,omitempty"`
	Name           string `json:"name,omitempty"`
	Decimals       int16  `json:"decimals,omitempty"`
	TotalSupply    BigInt `json:"totalSupply,omitempty"`
	DeploymentHash string `json:"deploymentHash,omitempty"`
}

type Tokens struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Token `json:"data,omitempty"`
}

type GetToken struct {
	Data Token `json:"data,omitempty"`
}

type TokenActionToken struct {
	Address  string `json:"address,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Decimals int16  `json:"decimals,omitempty"`
}

type TokenAction struct {
	TransactionHash string           `json:"transactionHash,omitempty"`
	From            string           `json:"from,omitempty"`
	To              string           `json:"to,omitempty"`
	Value           BigInt           `json:"value,omitempty"`
	FunctionSig     string           `json:"functionSig,omitempty"`
	BlockNumber     string           `json:"blockNumber,omitempty"`
	Timestamp       string           `json:"timestamp,omitempty"`
	Token           TokenActionToken `json:"token,omitempty"`
}

type TokenActions struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []TokenAction `json:"data,omitempty"`
}

type TokenActionsResults struct {
	Meta       Meta          `json:"meta,omitempty"`
	Results    []TokenAction `json:"results,omitempty"`
	TotalCount int64         `json:"totalCount,omitempty"`
}

type TokenHolder struct {
	TokenAddress string `json:"tokenAddress,omitempty"`
	Address      string `json:"address,omitempty"`
	Balance      BigInt `json:"balance,omitempty"`
}

type TokenHolders struct {
	Meta       Meta          `json:"meta,omitempty"`
	Results    []TokenHolder `json:"results,omitempty"`
	TotalCount int64         `json:"totalCount,omitempty"`
}

type TokenWhitelistEntry struct {
	Address   string `json:"address,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type TokenWhitelist struct {
	Meta Meta                  `json:"meta,omitempty"`
	Data []TokenWhitelistEntry `json:"data,omitempty"`
}

type WalletToken struct {
	TokenAddress string `json:"tokenAddress,omitempty"`
	Address      string `json:"address,omitempty"`
	Balance      BigInt `json:"balance,omitempty"`
	Name         string `json:"name,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	Decimals     int16  `json:"decimals,omitempty"`
	Supply       BigInt `json:"supply,omitempty"`
}

type WalletTokens struct {
	Meta Meta          `json:"meta,omitempty"`
	Data []WalletToken `json:"data,omitempty"`
}

type WalletTokenAddresses struct {
	Token     string            `json:"token,omitempty"`
	Symbol    string            `json:"symbol,omitempty"`
	Name      string            `json:"name,omitempty"`
	Decimals  int16             `json:"decimals,omitempty"`
	Supply    BigInt            `json:"supply,omitempty"`
	Addresses map[string]BigInt `json:"addresses,omitempty"`
}

type WalletTokenAddressesResponse struct {
	Meta Meta                   `json:"meta,omitempty"`
	Data []WalletTokenAddresses `json:"data,omitempty"`
}

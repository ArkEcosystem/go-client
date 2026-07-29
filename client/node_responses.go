// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type GetNodeStatus struct {
	Data NodeStatus `json:"data,omitempty"`
}

type GetNodeSyncing struct {
	Data NodeSyncing `json:"data,omitempty"`
}

type GetNodeConfiguration struct {
	Data NodeConfiguration `json:"data,omitempty"`
}

type GetNodeFees struct {
	Data NodeFeesResponse `json:"data,omitempty"`
}

type NodeStatus struct {
	Synced      bool  `json:"synced,omitempty"`
	Now         int64 `json:"now,omitempty"`
	BlocksCount int64 `json:"blocksCount,omitempty"`
	Timestamp   int64 `json:"timestamp,omitempty"`
}

type NodeSyncing struct {
	Syncing     bool  `json:"syncing,omitempty"`
	Blocks      int64 `json:"blocks,omitempty"`
	BlockNumber int64 `json:"blockNumber,omitempty"`
	Id          int   `json:"id,omitempty"`
}

type NodeCore struct {
	Version string `json:"version,omitempty"`
}

type NodeConfiguration struct {
	Nethash         string              `json:"nethash,omitempty"`
	Token           string              `json:"token,omitempty"`
	Symbol          string              `json:"symbol,omitempty"`
	Explorer        string              `json:"explorer,omitempty"`
	Version         int16               `json:"version,omitempty"`
	Wif             int16               `json:"wif,omitempty"`
	Slip44          int16               `json:"slip44,omitempty"`
	Ports           map[string]int16    `json:"ports,omitempty"`
	Constants       NodeConstants       `json:"constants,omitempty"`
	Core            NodeCore            `json:"core,omitempty"`
	TransactionPool NodeTransactionPool `json:"transactionPool,omitempty"`
}

type NodeConstantsGasLimits struct {
	Vote                  int64 `json:"vote,omitempty"`
	Transfer              int64 `json:"transfer,omitempty"`
	MultiPayment          int64 `json:"multiPayment,omitempty"`
	MultiSignature        int64 `json:"multiSignature,omitempty"`
	UsernameResignation   int64 `json:"usernameResignation,omitempty"`
	UsernameRegistration  int64 `json:"usernameRegistration,omitempty"`
	ValidatorResignation  int64 `json:"validatorResignation,omitempty"`
	ValidatorRegistration int64 `json:"validatorRegistration,omitempty"`
}

type NodeConstantsGas struct {
	MinimumGasFee       int64                  `json:"minimumGasFee,omitempty"`
	MaximumGasLimit     int64                  `json:"maximumGasLimit,omitempty"`
	MinimumGasLimit     int64                  `json:"minimumGasLimit,omitempty"`
	NativeGasLimits     NodeConstantsGasLimits `json:"nativeGasLimits,omitempty"`
	NativeFeeMultiplier int64                  `json:"nativeFeeMultiplier,omitempty"`
}

type NodeConstantsStaticFees struct {
	Vote                  int64 `json:"vote,omitempty"`
	Transfer              int64 `json:"transfer,omitempty"`
	MultiPayment          int64 `json:"multiPayment,omitempty"`
	MultiSignature        int64 `json:"multiSignature,omitempty"`
	UsernameResignation   int64 `json:"usernameResignation,omitempty"`
	UsernameRegistration  int64 `json:"usernameRegistration,omitempty"`
	ValidatorResignation  int64 `json:"validatorResignation,omitempty"`
	ValidatorRegistration int64 `json:"validatorRegistration,omitempty"`
}

type NodeConstantsFees struct {
	StaticFees NodeConstantsStaticFees `json:"staticFees,omitempty"`
}

type NodeConstantsBlock struct {
	Version         byte  `json:"version,omitempty"`
	MaxPayload      int64 `json:"maxPayload,omitempty"`
	MaxGasLimit     int64 `json:"maxGasLimit,omitempty"`
	MaxTransactions int64 `json:"maxTransactions,omitempty"`
}

type NodeConstantsAddress struct {
	Keccak256 bool `json:"keccak256,omitempty"`
}

type NodeConstantsSatoshi struct {
	Decimals     int16 `json:"decimals,omitempty"`
	Denomination int64 `json:"denomination,omitempty"`
}

type NodeConstantsTimeouts struct {
	BlockTime            int64 `json:"blockTime,omitempty"`
	Tolerance            int64 `json:"tolerance,omitempty"`
	StageTimeout         int64 `json:"stageTimeout,omitempty"`
	BlockPrepareTime     int64 `json:"blockPrepareTime,omitempty"`
	StageTimeoutIncrease int64 `json:"stageTimeoutIncrease,omitempty"`
}

type NodeConstants struct {
	Gas               NodeConstantsGas      `json:"gas,omitempty"`
	Fees              NodeConstantsFees     `json:"fees,omitempty"`
	Block             NodeConstantsBlock    `json:"block,omitempty"`
	Epoch             string                `json:"epoch,omitempty"`
	Height            int64                 `json:"height,omitempty"`
	Reward            BigInt                `json:"reward,omitempty"`
	Address           NodeConstantsAddress  `json:"address,omitempty"`
	EvmSpec           string                `json:"evmSpec,omitempty"`
	Satoshi           NodeConstantsSatoshi  `json:"satoshi,omitempty"`
	Timeouts          NodeConstantsTimeouts `json:"timeouts,omitempty"`
	ActiveValidators  int64                 `json:"activeValidators,omitempty"`
	MultiPaymentLimit int64                 `json:"multiPaymentLimit,omitempty"`
	VendorFieldLength int64                 `json:"vendorFieldLength,omitempty"`
}

type NodeDynamicFees struct {
	Enabled bool `json:"enabled,omitempty"`
}

type NodeTransactionPool struct {
	DynamicFees NodeDynamicFees `json:"dynamicFees,omitempty"`
}

// TransactionTypeFee represents the fee statistics for a single transaction type.
type TransactionTypeFee struct {
	Avg BigInt `json:"avg,omitempty"`
	Max BigInt `json:"max,omitempty"`
	Min BigInt `json:"min,omitempty"`
	Sum BigInt `json:"sum,omitempty"`
}

// NodeFeesResponse represents the response from the /node/fees endpoint, keyed by transaction type.
type NodeFeesResponse map[string]TransactionTypeFee

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
	Nethash   string           `json:"nethash,omitempty"`
	Token     string           `json:"token,omitempty"`
	Symbol    string           `json:"symbol,omitempty"`
	Explorer  string           `json:"explorer,omitempty"`
	Version   int16            `json:"version,omitempty"`
	Wif       int16            `json:"wif,omitempty"`
	Ports     map[string]int16 `json:"ports,omitempty"`
	Constants NodeConstants    `json:"constants,omitempty"`
	Core      NodeCore         `json:"core,omitempty"`
}

type NodeConstantsGas struct {
	MaximumGasLimit int64 `json:"maximumGasLimit,omitempty"`
	MaximumGasPrice int64 `json:"maximumGasPrice,omitempty"`
	MinimumGasLimit int64 `json:"minimumGasLimit,omitempty"`
	MinimumGasPrice int64 `json:"minimumGasPrice,omitempty"`
}

type NodeConstantsBlock struct {
	Version     byte  `json:"version,omitempty"`
	MaxPayload  int64 `json:"maxPayload,omitempty"`
	MaxGasLimit int64 `json:"maxGasLimit,omitempty"`
}

type NodeConstantsSatoshi struct {
	Decimals     int16 `json:"decimals,omitempty"`
	Denomination int64 `json:"denomination,omitempty"`
}

type NodeConstantsSnapshot struct {
	SnapshotHash             string `json:"snapshotHash,omitempty"`
	PreviousGenesisBlockHash string `json:"previousGenesisBlockHash,omitempty"`
}

type NodeConstantsTimeouts struct {
	BlockTime            int64 `json:"blockTime,omitempty"`
	Tolerance            int64 `json:"tolerance,omitempty"`
	StageTimeout         int64 `json:"stageTimeout,omitempty"`
	BlockPrepareTime     int64 `json:"blockPrepareTime,omitempty"`
	StageTimeoutIncrease int64 `json:"stageTimeoutIncrease,omitempty"`
}

type NodeConstants struct {
	Gas                      NodeConstantsGas      `json:"gas,omitempty"`
	Block                    NodeConstantsBlock    `json:"block,omitempty"`
	Epoch                    string                `json:"epoch,omitempty"`
	Height                   int64                 `json:"height,omitempty"`
	Reward                   BigInt                `json:"reward,omitempty"`
	EvmSpec                  string                `json:"evmSpec,omitempty"`
	Satoshi                  NodeConstantsSatoshi  `json:"satoshi,omitempty"`
	Snapshot                 NodeConstantsSnapshot `json:"snapshot,omitempty"`
	Timeouts                 NodeConstantsTimeouts `json:"timeouts,omitempty"`
	RoundValidators          int64                 `json:"roundValidators,omitempty"`
	ValidatorRegistrationFee BigInt                `json:"validatorRegistrationFee,omitempty"`
}

type TransactionTypeFee struct {
	Avg BigInt `json:"avg,omitempty"`
	Max BigInt `json:"max,omitempty"`
	Min BigInt `json:"min,omitempty"`
	Sum BigInt `json:"sum,omitempty"`
}

type NodeFeesResponse map[string]TransactionTypeFee

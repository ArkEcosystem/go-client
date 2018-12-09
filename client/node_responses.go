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

type NodeStatus struct {
	Synced      bool  `json:"synced,omitempty"`
	Now         int64 `json:"now,omitempty"`
	BlocksCount int   `json:"blocksCount,omitempty"`
}

type NodeSyncing struct {
	Syncing bool   `json:"syncing,omitempty"`
	Blocks  int    `json:"blocks,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Id      string `json:"id,omitempty"`
}

type NodeConfiguration struct {
	Nethash       string            `json:"nethash,omitempty"`
	Token         string            `json:"token,omitempty"`
	Symbol        string            `json:"symbol,omitempty"`
	Explorer      string            `json:"explorer,omitempty"`
	Version       int16             `json:"version,omitempty"`
	Ports         map[string]string `json:"ports,omitempty"`
	Constants     NodeConstants     `json:"constants,omitempty"`
	FeeStatistics []FeeStatistic    `json:"feeStatistics,omitempty"`
}

type NodeConstantsBlock struct {
	Version         byte  `json:"version,omitempty"`
	MaxTransactions byte  `json:"maxTransactions,omitempty"`
	MaxPayload      int64 `json:"maxPayload,omitempty"`
}

type NodeConstants struct {
	Height          int64              `json:"height,omitempty"`
	Reward          int64              `json:"reward,omitempty"`
	ActiveDelegates byte               `json:"activeDelegates,omitempty"`
	BlockTime       byte               `json:"blocktime,omitempty"`
	Block           NodeConstantsBlock `json:"block,omitempty"`
	Epoch           string             `json:"epoch,omitempty"`
	Fees            FeeTypes           `json:"fees,omitempty"`
	DynamicOffsets  DynamicFeeOffsets  `json:"dynamicOffsets,omitempty"`
}

type Fees struct {
	MinFee int64 `json:"minFee,omitempty"`
	MaxFee int64 `json:"maxFee,omitempty"`
	MvgFee int64 `json:"avgFee,omitempty"`
}

type FeeStatistic struct {
	Type int16 `json:"type,omitempty"`
	Fees Fees  `json:"fees,omitempty"`
}

type FeeTypes struct {
	Dynamic              bool  `json:"dynamic,omitempty"`
	Transfer             int64 `json:"transfer,omitempty"`
	SecondSignature      int64 `json:"secondSignature,omitempty"`
	DelegateRegistration int64 `json:"delegateRegistration,omitempty"`
	Vote                 int64 `json:"vote,omitempty"`
	MultiSignature       int64 `json:"multiSignature,omitempty"`
	Ipfs                 int64 `json:"ipfs,omitempty"`
	TimelockTransfer     int64 `json:"timelockTransfer,omitempty"`
	MultiPayment         int64 `json:"multiPayment,omitempty"`
	DelegateResignation  int64 `json:"delegateResignation,omitempty"`
}

type DynamicFeeOffsets struct {
	Transfer             int `json:"transfer,omitempty"`
	SecondSignature      int `json:"secondSignature,omitempty"`
	DelegateRegistration int `json:"delegateRegistration,omitempty"`
	Vote                 int `json:"vote,omitempty"`
	MultiSignature       int `json:"multiSignature,omitempty"`
	Ipfs                 int `json:"ipfs,omitempty"`
	TimelockTransfer     int `json:"timelockTransfer,omitempty"`
	MultiPayment         int `json:"multiPayment,omitempty"`
	DelegateResignation  int `json:"delegateResignation,omitempty"`
}

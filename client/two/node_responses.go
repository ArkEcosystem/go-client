// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

type NodeStatus struct {
	Success     bool  `json:"success,omitempty"`
	Synced      bool  `json:"synced,omitempty"`
	Now         int64 `json:"now,omitempty"`
	BlocksCount byte  `json:"blocksCount,omitempty"`
}

type NodeSyncing struct {
	Success bool   `json:"success,omitempty"`
	Syncing bool   `json:"syncing,omitempty"`
	Blocks  byte   `json:"blocks,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Id      string `json:"id,omitempty"`
}

type NodeConfiguration struct {
	Success       bool            `json:"success,omitempty"`
	Nethash       string          `json:"nethash,omitempty"`
	Token         string          `json:"token,omitempty"`
	Symbol        string          `json:"symbol,omitempty"`
	Explorer      string          `json:"explorer,omitempty"`
	Version       string          `json:"version,omitempty"`
	Ports         map[string]byte `json:"ports,omitempty"`
	Constants     []NodeConstants `json:"constants,omitempty"`
	FeeStatistics []FeeStatistic  `json:"feeStatistics,omitempty"`
}

type NodeConstantsBlock struct {
	Version         byte `json:"version,omitempty"`
	MaxTransactions byte `json:"maxTransactions,omitempty"`
	MaxPayload      byte `json:"maxPayload,omitempty"`
}

type NodeConstants struct {
	Height          int64              `json:"height,omitempty"`
	Reward          int64              `json:"reward,omitempty"`
	ActiveDelegates byte               `json:"activeDelegates,omitempty"`
	BlockTime       byte               `json:"blocktime,omitempty"`
	Block           NodeConstantsBlock `json:"block,omitempty"`
	Epoch           byte               `json:"epoch,omitempty"`
	Fees            map[string]byte    `json:"fees,omitempty"`
	DynamicOffsets  map[string]byte    `json:"dynamicOffsets,omitempty"`
}

type Fees struct {
	MinFee int64 `json:"minFee,omitempty"`
	MaxFee int64 `json:"maxFee,omitempty"`
	MvgFee int64 `json:"avgFee,omitempty"`
}

type FeeStatistic struct {
	Type int64 `json:"type,omitempty"`
	Fees Fees  `json:"fees,omitempty"`
}

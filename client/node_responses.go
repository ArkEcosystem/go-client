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
	Data NodeFees `json:"data,omitempty"`
}

type NodeStatus struct {
	Synced      bool  `json:"synced,omitempty"`
	Now         int64 `json:"now,omitempty"`
	BlocksCount int64 `json:"blocksCount,omitempty"`
	Timestamp   int64 `json:"timestamp,omitempty"`
}

type NodeSyncing struct {
	Syncing bool   `json:"syncing,omitempty"`
	Blocks  int64  `json:"blocks,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Id      string `json:"id,omitempty"`
}

type NodeConfiguration struct {
	Nethash         string              `json:"nethash,omitempty"`
	Token           string              `json:"token,omitempty"`
	Symbol          string              `json:"symbol,omitempty"`
	Explorer        string              `json:"explorer,omitempty"`
	Version         int16               `json:"version,omitempty"`
	Ports           map[string]int16    `json:"ports,omitempty"`
	Constants       NodeConstants       `json:"constants,omitempty"`
	TransactionPool NodeTransactionPool `json:"transactionPool,omitempty"`
}

type NodeFees []FeeStatistic

type NodeConstantsBlock struct {
	Version         byte  `json:"version,omitempty"`
	MaxTransactions byte  `json:"maxTransactions,omitempty"`
	MaxPayload      int64 `json:"maxPayload,omitempty"`
}

type NodeConstants struct {
	Height          int64               `json:"height,omitempty"`
	Reward          int64               `json:"reward,omitempty"`
	ActiveDelegates byte                `json:"activeDelegates,omitempty"`
	BlockTime       byte                `json:"blocktime,omitempty"`
	Block           NodeConstantsBlock  `json:"block,omitempty"`
	Epoch           string              `json:"epoch,omitempty"`
	Fees            map[string]FeeTypes `json:"fees,omitempty"`
}

type DynamicFees struct {
	Enabled         bool     `json:"enabled,omitempty"`
	MinFeePool      int16    `json:"minFeePool,omitempty"`
	MinFeeBroadcast int16    `json:"minFeeBroadcast,omitempty"`
	AddonBytes      FeeTypes `json:"addonBytes,omitempty"`
}

type NodeTransactionPool struct {
	DynamicFees DynamicFees `json:"dynamicFees,omitempty"`
}

type FeeStatistic struct {
	Type   int16  `json:"type,omitempty"`
	MinFee string `json:"min,omitempty"`
	MaxFee string `json:"max,omitempty"`
	AvgFee string `json:"avg,omitempty"`
	SumFee string `json:"sum,omitempty"`
	MdnFee string `json:"median,omitempty"`
}

type FeeTypes struct {
	Transfer             int64 `json:"transfer,omitempty"`
	SecondSignature      int64 `json:"secondSignature,omitempty"`
	DelegateRegistration int64 `json:"delegateRegistration,omitempty"`
	Vote                 int64 `json:"vote,omitempty"`
	MultiSignature       int64 `json:"multiSignature,omitempty"`
	Ipfs                 int64 `json:"ipfs,omitempty"`
	MultiPayment         int64 `json:"multiPayment,omitempty"`
	DelegateResignation  int64 `json:"delegateResignation,omitempty"`
	HtlcLock             int64 `json:"htlcLock,omitempty"`
	HtlcClaim            int64 `json:"htlcClaim,omitempty"`
	HtlcRefund           int64 `json:"htlcRefund,omitempty"`
}

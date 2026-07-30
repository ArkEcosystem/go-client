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

// NodeConstants represents a single milestone's constants, as returned both
// by node/configuration (the currently active milestone) and as an entry in
// node/configuration/crypto's milestones list.
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

type GetNodeCrypto struct {
	Data NodeCrypto `json:"data,omitempty"`
}

type NetworkClient struct {
	Token    string `json:"token,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Explorer string `json:"explorer,omitempty"`
}

type Network struct {
	Wif        int16         `json:"wif,omitempty"`
	Name       string        `json:"name,omitempty"`
	Client     NetworkClient `json:"client,omitempty"`
	ChainId    int64         `json:"chainId,omitempty"`
	Nethash    string        `json:"nethash,omitempty"`
	PubKeyHash int16         `json:"pubKeyHash,omitempty"`
}

// GenesisBlock represents the genesis block as returned by
// node/configuration/crypto. It intentionally does not reuse Block: several
// fields differ (e.g. Timestamp is a raw number here, a string on Block),
// and it carries fields Block doesn't (logsBloom, serialized).
type GenesisBlock struct {
	Fee               BigInt        `json:"fee,omitempty"`
	Hash              string        `json:"hash,omitempty"`
	Round             int64         `json:"round,omitempty"`
	Number            int64         `json:"number,omitempty"`
	Reward            BigInt        `json:"reward,omitempty"`
	GasUsed           int64         `json:"gasUsed,omitempty"`
	Version           byte          `json:"version,omitempty"`
	Proposer          string        `json:"proposer,omitempty"`
	LogsBloom         string        `json:"logsBloom,omitempty"`
	StateRoot         string        `json:"stateRoot,omitempty"`
	Timestamp         int64         `json:"timestamp,omitempty"`
	ParentHash        string        `json:"parentHash,omitempty"`
	Serialized        string        `json:"serialized,omitempty"`
	PayloadSize       int64         `json:"payloadSize,omitempty"`
	Transactions      []interface{} `json:"transactions,omitempty"`
	TransactionsRoot  string        `json:"transactionsRoot,omitempty"`
	TransactionsCount int64         `json:"transactionsCount,omitempty"`
}

type GenesisBlockProof struct {
	Round      int64    `json:"round,omitempty"`
	Signature  string   `json:"signature,omitempty"`
	Validators []string `json:"validators,omitempty"`
}

type NodeCryptoGenesisBlock struct {
	Block      GenesisBlock      `json:"block,omitempty"`
	Proof      GenesisBlockProof `json:"proof,omitempty"`
	Serialized string            `json:"serialized,omitempty"`
}

type NodeCrypto struct {
	Network      Network                `json:"network,omitempty"`
	Milestones   []NodeConstants        `json:"milestones,omitempty"`
	GenesisBlock NodeCryptoGenesisBlock `json:"genesisBlock,omitempty"`
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

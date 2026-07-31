// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get the node status.
func TestNodeService_Status(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/node/status", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "synced": true,
			    "now": 2399684,
			    "blocksCount": -1346,
			    "timestamp": 82359359
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Status(context.Background())
	testGeneralError(t, "Node.Status", err)
	testResponseUrl(t, "Node.Status", response, "/api/node/status")
	testResponseStruct(t, "Node.Status", responseStruct, &GetNodeStatus{
		Data: NodeStatus{
			Synced:      true,
			Now:         2399684,
			BlocksCount: -1346,
			Timestamp:   82359359,
		},
	})
}

// Get the node syncing status.
func TestNodeService_Syncing(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/node/syncing", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "syncing": false,
			    "blocks": -1385,
			    "blockNumber": 2399723,
			    "id": 1
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Syncing(context.Background())
	testGeneralError(t, "Node.Syncing", err)
	testResponseUrl(t, "Node.Syncing", response, "/api/node/syncing")
	testResponseStruct(t, "Node.Syncing", responseStruct, &GetNodeSyncing{
		Data: NodeSyncing{
			Syncing:     false,
			Blocks:      -1385,
			BlockNumber: 2399723,
			Id:          1,
		},
	})
}

// Get the node configuration.
func TestNodeService_Configuration(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/node/configuration", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "constants": {
			      "gas": {
			        "maximumGasLimit": 5000000,
			        "maximumGasPrice": 10000000000000,
			        "minimumGasLimit": 21000,
			        "minimumGasPrice": 5000000000
			      },
			      "block": {
			        "version": 1,
			        "maxPayload": 2097152,
			        "maxGasLimit": 10000000
			      },
			      "epoch": "2026-06-02T00:00:00.000Z",
			      "height": 22839038,
			      "reward": "2000000000000000000",
			      "evmSpec": "Shanghai",
			      "satoshi": {
			        "decimals": 18,
			        "denomination": 1000000000000000000
			      },
			      "snapshot": {
			        "snapshotHash": "002d1b892ca9c970c53c8576907971593cd576749b4b26ea8ec7f3f52688aed7",
			        "previousGenesisBlockHash": "c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7"
			      },
			      "timeouts": {
			        "blockTime": 8000,
			        "tolerance": 500,
			        "stageTimeout": 2000,
			        "blockPrepareTime": 4000,
			        "stageTimeoutIncrease": 2000
			      },
			      "roundValidators": 53,
			      "validatorRegistrationFee": "250000000000000000000"
			    },
			    "core": {
			      "version": "0.0.1-evm.53"
			    },
			    "explorer": "",
			    "nethash": "560f869ed6713745a12328e7214cb65077e645bb5e57b1e5b323bb915a51f114",
			    "ports": {
			      "@mainsail/api-database": null
			    },
			    "symbol": "TѦ",
			    "token": "ARK",
			    "version": 30,
			    "wif": 186
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Configuration(context.Background())
	testGeneralError(t, "Node.Configuration", err)
	testResponseUrl(t, "Node.Configuration", response, "/api/node/configuration")
	testResponseStruct(t, "Node.Configuration", responseStruct, &GetNodeConfiguration{
		Data: NodeConfiguration{
			Nethash:  "560f869ed6713745a12328e7214cb65077e645bb5e57b1e5b323bb915a51f114",
			Token:    "ARK",
			Symbol:   "TѦ",
			Explorer: "",
			Version:  30,
			Wif:      186,
			Ports: map[string]int64{
				"@mainsail/api-database": 0,
			},
			Core: NodeCore{
				Version: "0.0.1-evm.53",
			},
			Constants: NodeConstants{
				Gas: NodeConstantsGas{
					MaximumGasLimit: 5000000,
					MaximumGasPrice: 10000000000000,
					MinimumGasLimit: 21000,
					MinimumGasPrice: 5000000000,
				},
				Block: NodeConstantsBlock{
					Version:     1,
					MaxPayload:  2097152,
					MaxGasLimit: 10000000,
				},
				Epoch:   "2026-06-02T00:00:00.000Z",
				Height:  22839038,
				Reward:  newBigInt(2000000000000000000),
				EvmSpec: "Shanghai",
				Satoshi: NodeConstantsSatoshi{
					Decimals:     18,
					Denomination: 1000000000000000000,
				},
				Snapshot: NodeConstantsSnapshot{
					SnapshotHash:             "002d1b892ca9c970c53c8576907971593cd576749b4b26ea8ec7f3f52688aed7",
					PreviousGenesisBlockHash: "c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7",
				},
				Timeouts: NodeConstantsTimeouts{
					BlockTime:            8000,
					Tolerance:            500,
					StageTimeout:         2000,
					BlockPrepareTime:     4000,
					StageTimeoutIncrease: 2000,
				},
				RoundValidators:          53,
				ValidatorRegistrationFee: newBigIntFromString("250000000000000000000"),
			},
		},
	})
}

func TestNodeService_Crypto(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/node/configuration/crypto", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "network": {
			      "wif": 186,
			      "name": "testnet",
			      "client": {"token": "ARK", "symbol": "TѦ", "explorer": ""},
			      "chainId": 11812,
			      "nethash": "560f869ed6713745a12328e7214cb65077e645bb5e57b1e5b323bb915a51f114",
			      "pubKeyHash": 30
			    },
			    "milestones": [
			      {
			        "gas": {"maximumGasLimit":5000000,"maximumGasPrice":10000000000000,"minimumGasLimit":21000,"minimumGasPrice":5000000000},
			        "block": {"version":1,"maxPayload":2097152,"maxGasLimit":10000000},
			        "epoch": "2026-06-02T00:00:00.000Z",
			        "height": 22763438,
			        "reward": "0",
			        "evmSpec": "Shanghai",
			        "satoshi": {"decimals":18,"denomination":1000000000000000000},
			        "snapshot": {"snapshotHash":"002d1b892ca9c970c53c8576907971593cd576749b4b26ea8ec7f3f52688aed7","previousGenesisBlockHash":"c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7"},
			        "timeouts": {"blockTime":8000,"tolerance":500,"stageTimeout":2000,"blockPrepareTime":4000,"stageTimeoutIncrease":2000},
			        "roundValidators": 0,
			        "validatorRegistrationFee": "250000000000000000000"
			      }
			    ],
			    "genesisBlock": {
			      "block": {
			        "fee":"0","hash":"7e0a16b27cc23a7614f35a15a33b112e3f992567afe90b32e1c7692ada9e815f","round":0,"number":22763438,"reward":"0","gasUsed":0,"version":1,
			        "proposer":"0x8065f3ff4879d6423d9E5693b8c3AFAE6B806b28",
			        "logsBloom":"00",
			        "stateRoot":"6d81ecfc4e380bfefee8c80c8e178335a09b8c81ad83d52317dab7ff4d7f1c3f","timestamp":1780373265937,"parentHash":"c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7",
			        "serialized":"01111e84",
			        "payloadSize":0,"transactions":[],"transactionsRoot":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","transactionsCount":0
			      },
			      "proof": {"round":0,"signature":"00","validators":[]},
			      "serialized":"00"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Crypto(context.Background())
	testGeneralError(t, "Node.Crypto", err)
	testResponseUrl(t, "Node.Crypto", response, "/api/node/configuration/crypto")
	testResponseStruct(t, "Node.Crypto", responseStruct, &GetNodeCrypto{
		Data: NodeCrypto{
			Network: Network{
				Wif:  186,
				Name: "testnet",
				Client: NetworkClient{
					Token:    "ARK",
					Symbol:   "TѦ",
					Explorer: "",
				},
				ChainId:    11812,
				Nethash:    "560f869ed6713745a12328e7214cb65077e645bb5e57b1e5b323bb915a51f114",
				PubKeyHash: 30,
			},
			Milestones: []NodeConstants{{
				Gas: NodeConstantsGas{
					MaximumGasLimit: 5000000,
					MaximumGasPrice: 10000000000000,
					MinimumGasLimit: 21000,
					MinimumGasPrice: 5000000000,
				},
				Block: NodeConstantsBlock{
					Version:     1,
					MaxPayload:  2097152,
					MaxGasLimit: 10000000,
				},
				Epoch:   "2026-06-02T00:00:00.000Z",
				Height:  22763438,
				Reward:  newBigInt(0),
				EvmSpec: "Shanghai",
				Satoshi: NodeConstantsSatoshi{
					Decimals:     18,
					Denomination: 1000000000000000000,
				},
				Snapshot: NodeConstantsSnapshot{
					SnapshotHash:             "002d1b892ca9c970c53c8576907971593cd576749b4b26ea8ec7f3f52688aed7",
					PreviousGenesisBlockHash: "c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7",
				},
				Timeouts: NodeConstantsTimeouts{
					BlockTime:            8000,
					Tolerance:            500,
					StageTimeout:         2000,
					BlockPrepareTime:     4000,
					StageTimeoutIncrease: 2000,
				},
				RoundValidators:          0,
				ValidatorRegistrationFee: newBigIntFromString("250000000000000000000"),
			}},
			GenesisBlock: NodeCryptoGenesisBlock{
				Block: GenesisBlock{
					Fee:               newBigInt(0),
					Hash:              "7e0a16b27cc23a7614f35a15a33b112e3f992567afe90b32e1c7692ada9e815f",
					Round:             0,
					Number:            22763438,
					Reward:            newBigInt(0),
					GasUsed:           0,
					Version:           1,
					Proposer:          "0x8065f3ff4879d6423d9E5693b8c3AFAE6B806b28",
					LogsBloom:         "00",
					StateRoot:         "6d81ecfc4e380bfefee8c80c8e178335a09b8c81ad83d52317dab7ff4d7f1c3f",
					Timestamp:         1780373265937,
					ParentHash:        "c37ec049e0317eb8f03db8bcfc5911551f89ba132042ce8a583b1ef0590806f7",
					Serialized:        "01111e84",
					PayloadSize:       0,
					Transactions:      []interface{}{},
					TransactionsRoot:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
					TransactionsCount: 0,
				},
				Proof: GenesisBlockProof{
					Round:      0,
					Signature:  "00",
					Validators: []string{},
				},
				Serialized: "00",
			},
		},
	})
}

// Get the node fee statistics.
func TestNodeService_Fees(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/node/fees", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "evmCall": {
			      "avg": "8480594889",
			      "max": "8888977300",
			      "min": "5000000000",
			      "sum": "169611897796"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Fees(context.Background(), 7)
	testGeneralError(t, "Node.Fees", err)
	testResponseUrl(t, "Node.Fees", response, "/api/node/fees?days=7")
	testResponseStruct(t, "Node.Fees", responseStruct, &GetNodeFees{
		Data: NodeFeesResponse{
			"evmCall": {
				Avg: newBigInt(8480594889),
				Max: newBigInt(8888977300),
				Min: newBigInt(5000000000),
				Sum: newBigInt(169611897796),
			},
		},
	})
}

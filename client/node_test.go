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
			Ports: map[string]int16{
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

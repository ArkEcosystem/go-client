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
			    "nethash": "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23",
			    "token": "DARK",
			    "symbol": "DѦ",
			    "explorer": "https://dwallets.ark.io/api/",
			    "version": 30,
			    "wif": 170,
			    "slip44": 111,
			    "ports": {
			      "@arkecosystem/core-p2p": 4002,
			      "@arkecosystem/core-api": 4003
			    },
			    "core": {
			      "version": "4.0.0"
			    },
			    "constants": {
			      "height": 75600,
			      "reward": "200000000",
			      "activeValidators": 51,
			      "multiPaymentLimit": 128,
			      "vendorFieldLength": 255,
			      "epoch": "2017-03-21T13:00:00.000Z",
			      "evmSpec": "shanghai",
			      "block": {
			        "version": 0,
			        "maxPayload": 2097152,
			        "maxGasLimit": 30000000,
			        "maxTransactions": 50
			      },
			      "address": {
			        "keccak256": true
			      },
			      "satoshi": {
			        "decimals": 18,
			        "denomination": 1000000000000000000
			      },
			      "timeouts": {
			        "blockTime": 8000,
			        "tolerance": 500,
			        "stageTimeout": 2000,
			        "blockPrepareTime": 4000,
			        "stageTimeoutIncrease": 2000
			      },
			      "gas": {
			        "minimumGasFee": 10,
			        "maximumGasLimit": 30000000,
			        "minimumGasLimit": 21000,
			        "nativeFeeMultiplier": 1,
			        "nativeGasLimits": {
			          "vote": 100000,
			          "transfer": 21000,
			          "multiPayment": 50000,
			          "multiSignature": 60000,
			          "usernameResignation": 40000,
			          "usernameRegistration": 40000,
			          "validatorResignation": 40000,
			          "validatorRegistration": 60000
			        }
			      },
			      "fees": {
			        "staticFees": {
			          "vote": 100000000,
			          "transfer": 10000000,
			          "multiPayment": 10000000,
			          "multiSignature": 500000000,
			          "usernameResignation": 2500000000,
			          "usernameRegistration": 2500000000,
			          "validatorResignation": 2500000000,
			          "validatorRegistration": 2500000000
			        }
			      }
			    },
			    "transactionPool": {
			      "dynamicFees": {
			        "enabled": true
			      }
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Configuration(context.Background())
	testGeneralError(t, "Node.Configuration", err)
	testResponseUrl(t, "Node.Configuration", response, "/api/node/configuration")
	testResponseStruct(t, "Node.Configuration", responseStruct, &GetNodeConfiguration{
		Data: NodeConfiguration{
			Nethash:  "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23",
			Token:    "DARK",
			Symbol:   "DѦ",
			Explorer: "https://dwallets.ark.io/api/",
			Version:  30,
			Wif:      170,
			Slip44:   111,
			Ports: map[string]int16{
				"@arkecosystem/core-p2p": 4002,
				"@arkecosystem/core-api": 4003,
			},
			Core: NodeCore{
				Version: "4.0.0",
			},
			Constants: NodeConstants{
				Height:            75600,
				Reward:            newBigInt(200000000),
				ActiveValidators:  51,
				MultiPaymentLimit: 128,
				VendorFieldLength: 255,
				Epoch:             "2017-03-21T13:00:00.000Z",
				EvmSpec:           "shanghai",
				Block: NodeConstantsBlock{
					Version:         0,
					MaxPayload:      2097152,
					MaxGasLimit:     30000000,
					MaxTransactions: 50,
				},
				Address: NodeConstantsAddress{
					Keccak256: true,
				},
				Satoshi: NodeConstantsSatoshi{
					Decimals:     18,
					Denomination: 1000000000000000000,
				},
				Timeouts: NodeConstantsTimeouts{
					BlockTime:            8000,
					Tolerance:            500,
					StageTimeout:         2000,
					BlockPrepareTime:     4000,
					StageTimeoutIncrease: 2000,
				},
				Gas: NodeConstantsGas{
					MinimumGasFee:       10,
					MaximumGasLimit:     30000000,
					MinimumGasLimit:     21000,
					NativeFeeMultiplier: 1,
					NativeGasLimits: NodeConstantsGasLimits{
						Vote:                  100000,
						Transfer:              21000,
						MultiPayment:          50000,
						MultiSignature:        60000,
						UsernameResignation:   40000,
						UsernameRegistration:  40000,
						ValidatorResignation:  40000,
						ValidatorRegistration: 60000,
					},
				},
				Fees: NodeConstantsFees{
					StaticFees: NodeConstantsStaticFees{
						Vote:                  100000000,
						Transfer:              10000000,
						MultiPayment:          10000000,
						MultiSignature:        500000000,
						UsernameResignation:   2500000000,
						UsernameRegistration:  2500000000,
						ValidatorResignation:  2500000000,
						ValidatorRegistration: 2500000000,
					},
				},
			},
			TransactionPool: NodeTransactionPool{
				DynamicFees: NodeDynamicFees{
					Enabled: true,
				},
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
			    "transfer": {
			      "min": "10000000",
			      "max": "10000000",
			      "avg": "10000000",
			      "sum": "10000000"
			    },
			    "multiSignature": {
			      "min": "500000000",
			      "max": "500000000",
			      "avg": "500000000",
			      "sum": "500000000"
			    },
			    "vote": {
			      "min": "100000000",
			      "max": "100000000",
			      "avg": "100000000",
			      "sum": "100000000"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Fees(context.Background(), 7)
	testGeneralError(t, "Node.Fees", err)
	testResponseUrl(t, "Node.Fees", response, "/api/node/fees?days=7")
	testResponseStruct(t, "Node.Fees", responseStruct, &GetNodeFees{
		Data: NodeFeesResponse{
			"transfer": {
				Min: newBigInt(10000000),
				Max: newBigInt(10000000),
				Avg: newBigInt(10000000),
				Sum: newBigInt(10000000),
			},
			"multiSignature": {
				Min: newBigInt(500000000),
				Max: newBigInt(500000000),
				Avg: newBigInt(500000000),
				Sum: newBigInt(500000000),
			},
			"vote": {
				Min: newBigInt(100000000),
				Max: newBigInt(100000000),
				Avg: newBigInt(100000000),
				Sum: newBigInt(100000000),
			},
		},
	})
}

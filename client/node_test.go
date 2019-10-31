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
			    "height": 2399723,
			    "id": "10438786023074296467"
			  }
			}`)
	})

	responseStruct, response, err := client.Node.Syncing(context.Background())
	testGeneralError(t, "Node.Syncing", err)
	testResponseUrl(t, "Node.Syncing", response, "/api/node/syncing")
	testResponseStruct(t, "Node.Syncing", responseStruct, &GetNodeSyncing{
		Data: NodeSyncing{
			Syncing: false,
			Blocks:  -1385,
			Height:  2399723,
			Id:      "10438786023074296467",
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
			    "explorer": "https://dexplorer.ark.io",
			    "version": 30,
			    "ports": {
			      "@arkecosystem/core-p2p": 4002,
			      "@arkecosystem/core-api": 4003
			    },
			    "constants": {
			      "height": 75600,
			      "reward": 200000000,
			      "activeDelegates": 51,
			      "blocktime": 8,
			      "block": {
			        "version": 0,
			        "maxTransactions": 50,
			        "maxPayload": 2097152
			      },
			      "epoch": "2017-03-21T13:00:00.000Z",
			      "fees": {
			        "staticFees": {
			          "transfer": 10000000,
			          "secondSignature": 500000000,
			          "delegateRegistration": 2500000000,
			          "vote": 100000000,
			          "multiSignature": 500000000,
			          "ipfs": 500000000,
			          "multiPayment": 10000000,
			          "delegateResignation": 2500000000,
			          "htlcLock": 10000000,
			          "htlcClaim": 0,
			          "htlcRefund": 0
			        }
			      }
			    },
			    "transactionPool": {
			      "dynamicFees": {
			        "enabled": true,
			        "minFeePool": 1000,
			        "minFeeBroadcast": 1000,
			        "addonBytes": {
			          "transfer": 100,
			          "secondSignature": 250,
			          "delegateRegistration": 400000,
			          "vote": 100,
			          "multiSignature": 500,
			          "ipfs": 250,
			          "multiPayment": 500,
			          "delegateResignation": 100,
			          "htlcLock": 100,
			          "htlcClaim": 0,
			          "htlcRefund": 0
			        }
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
			Explorer: "https://dexplorer.ark.io",
			Version:  30,
			Ports: map[string]int16{
				"@arkecosystem/core-p2p": 4002,
				"@arkecosystem/core-api": 4003,
			},
			Constants: NodeConstants{
				Height:          75600,
				Reward:          200000000,
				ActiveDelegates: 51,
				BlockTime:       8,
				Block: NodeConstantsBlock{
					Version:         0,
					MaxTransactions: 50,
					MaxPayload:      2097152,
				},
				Epoch: "2017-03-21T13:00:00.000Z",
				Fees: map[string]FeeTypes{
					"staticFees": FeeTypes{
						Transfer:             10000000,
						SecondSignature:      500000000,
						DelegateRegistration: 2500000000,
						Vote:                 100000000,
						MultiSignature:       500000000,
						Ipfs:                 500000000,
						MultiPayment:         10000000,
						DelegateResignation:  2500000000,
						HtlcLock:             10000000,
						HtlcClaim:            0,
						HtlcRefund:           0,
					},
				},
			},
			TransactionPool: NodeTransactionPool{
				DynamicFees: DynamicFees{
					Enabled:         true,
					MinFeePool:      1000,
					MinFeeBroadcast: 1000,
					AddonBytes: FeeTypes{
						Transfer:             100,
						SecondSignature:      250,
						DelegateRegistration: 400000,
						Vote:                 100,
						MultiSignature:       500,
						Ipfs:                 250,
						MultiPayment:         500,
						DelegateResignation:  100,
						HtlcLock:             100,
						HtlcClaim:            0,
						HtlcRefund:           0,
					},
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
			  "data": [
			    {
			      "type": 0,
			      "min": "10000000",
			      "max": "10000000",
			      "avg": "10000000",
			      "sum": "10000000",
			      "median": "10000000"
			    },
			    {
			      "type": 1,
			      "min": "500000000",
			      "max": "500000000",
			      "avg": "500000000",
			      "sum": "500000000",
			      "median": "500000000"
			    },
			    {
			      "type": 3,
			      "min": "100000000",
			      "max": "100000000",
			      "avg": "100000000",
			      "sum": "100000000",
			      "median": "100000000"
			    }
			  ]
			}`)
	})

	responseStruct, response, err := client.Node.Fees(context.Background(), 7)
	testGeneralError(t, "Node.Fees", err)
	testResponseUrl(t, "Node.Fees", response, "/api/node/fees?days=7")
	testResponseStruct(t, "Node.Fees", responseStruct, &GetNodeFees{
		Data: []FeeStatistic{{
			Type:   0,
			MinFee: 10000000,
			MaxFee: 10000000,
			AvgFee: 10000000,
			SumFee: 10000000,
			MdnFee: 10000000,
		}, {
			Type:   1,
			MinFee: 500000000,
			MaxFee: 500000000,
			AvgFee: 500000000,
			SumFee: 500000000,
			MdnFee: 500000000,
		}, {
			Type:   3,
			MinFee: 100000000,
			MaxFee: 100000000,
			AvgFee: 100000000,
			SumFee: 100000000,
			MdnFee: 100000000,
		}},
	})
}

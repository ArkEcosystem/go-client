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
			    "blocksCount": -1346
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
			      "@arkecosystem/core-p2p": "4002",
			      "@arkecosystem/core-api": "4003"
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
			        "dynamic": false,
			        "transfer": 10000000,
			        "secondSignature": 500000000,
			        "delegateRegistration": 2500000000,
			        "vote": 100000000,
			        "multiSignature": 500000000,
			        "ipfs": 0,
			        "timelockTransfer": 0,
			        "multiPayment": 0,
			        "delegateResignation": 0
			      },
			      "dynamicOffsets": {
			        "transfer": 100,
			        "secondSignature": 250,
			        "delegateRegistration": 500,
			        "vote": 100,
			        "multiSignature": 500,
			        "ipfs": 250,
			        "timelockTransfer": 500,
			        "multiPayment": 500,
			        "delegateResignation": 500
			      }
			    },
			    "feeStatistics": [
			      {
			        "type": 0,
			        "fees": {
			          "minFee": 10000000,
			          "maxFee": 10000000,
			          "avgFee": 10000000
			        }
			      },
			      {
			        "type": 1,
			        "fees": {
			          "minFee": 500000000,
			          "maxFee": 500000000,
			          "avgFee": 500000000
			        }
			      },
			      {
			        "type": 3,
			        "fees": {
			          "minFee": 100000000,
			          "maxFee": 100000000,
			          "avgFee": 100000000
			        }
			      }
			    ]
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
			Ports: map[string]string{
				"@arkecosystem/core-p2p": "4002",
				"@arkecosystem/core-api": "4003",
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
				Fees: FeeTypes{
					Dynamic:              false,
					Transfer:             10000000,
					SecondSignature:      500000000,
					DelegateRegistration: 2500000000,
					Vote:                 100000000,
					MultiSignature:       500000000,
					Ipfs:                 0,
					TimelockTransfer:     0,
					MultiPayment:         0,
					DelegateResignation:  0,
				},
				DynamicOffsets: DynamicFeeOffsets{
					Transfer:             100,
					SecondSignature:      250,
					DelegateRegistration: 500,
					Vote:                 100,
					MultiSignature:       500,
					Ipfs:                 250,
					TimelockTransfer:     500,
					MultiPayment:         500,
					DelegateResignation:  500,
				},
			},
			FeeStatistics: []FeeStatistic{{
				Type: 0,
				Fees: Fees{
					MinFee: 10000000,
					MaxFee: 10000000,
					MvgFee: 10000000,
				},
			}, {
				Type: 1,
				Fees: Fees{
					MinFee: 500000000,
					MaxFee: 500000000,
					MvgFee: 500000000,
				},
			}, {
				Type: 3,
				Fees: Fees{
					MinFee: 100000000,
					MaxFee: 100000000,
					MvgFee: 100000000,
				},
			}},
		},
	})
}

// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all accounts.
func TestBlocksService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "blocks": [
			    {
			      "id": "dummy",
			      "version": 10,
			      "timestamp": 10,
			      "height": 10,
			      "previousBlock": "dummy",
			      "numberOfTransactions": 10,
			      "totalAmount": 10,
			      "totalFee": 10,
			      "reward": 10,
			      "payloadLength": 10,
			      "payloadHash": "dummy",
			      "generatorPublicKey": "dummy",
			      "generatorId": "dummy",
			      "blockSignature": "dummy",
			      "confirmations": 10,
			      "totalForged": "10"
			    }
			  ]
			}`)
	})

	query := &BlocksQuery{Limit: 1}
	responseStruct, response, err := client.Blocks.List(context.Background(), query)
	testGeneralError(t, "Blocks.List", err)
	testResponseUrl(t, "Blocks.List", response, "/api/blocks")
	testResponseStruct(t, "Blocks.List", responseStruct, &Blocks{
		Success: true,
		Blocks: []Block{{
			Id:                   "dummy",
			Version:              10,
			Timestamp:            10,
			Height:               10,
			PreviousBlock:        "dummy",
			NumberOfTransactions: 10,
			TotalAmount:          10,
			TotalFee:             10,
			Reward:               10,
			PayloadLength:        10,
			PayloadHash:          "dummy",
			GeneratorPublicKey:   "dummy",
			GeneratorId:          "dummy",
			BlockSignature:       "dummy",
			Confirmations:        10,
			TotalForged:          "10",
		}},
	})
}

// Get a block by the given id.
func TestBlocksService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/get", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "block": {
			    "id": "dummy",
			    "version": 10,
			    "timestamp": 10,
			    "height": 10,
			    "previousBlock": "dummy",
			    "numberOfTransactions": 10,
			    "totalAmount": 10,
			    "totalFee": 10,
			    "reward": 10,
			    "payloadLength": 10,
			    "payloadHash": "dummy",
			    "generatorPublicKey": "dummy",
			    "generatorId": "dummy",
			    "blockSignature": "dummy",
			    "confirmations": 10,
			    "totalForged": "10"
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Get(context.Background(), "dummy")
	testGeneralError(t, "Blocks.Get", err)
	testResponseUrl(t, "Blocks.Get", response, "/api/blocks/get")
	testResponseStruct(t, "Blocks.Get", responseStruct, &GetBlock{
		Success: true,
		Block: Block{
			Id:                   "dummy",
			Version:              10,
			Timestamp:            10,
			Height:               10,
			PreviousBlock:        "dummy",
			NumberOfTransactions: 10,
			TotalAmount:          10,
			TotalFee:             10,
			Reward:               10,
			PayloadLength:        10,
			PayloadHash:          "dummy",
			GeneratorPublicKey:   "dummy",
			GeneratorId:          "dummy",
			BlockSignature:       "dummy",
			Confirmations:        10,
			TotalForged:          "10",
		},
	})
}

// Get the blockchain epoch.
func TestBlocksService_Epoch(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getEpoch", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "epoch": "2017-03-21T13:00:00.000Z"
			}`)
	})

	responseStruct, response, err := client.Blocks.Epoch(context.Background())
	testGeneralError(t, "Blocks.Epoch", err)
	testResponseUrl(t, "Blocks.Epoch", response, "/api/blocks/getEpoch")
	testResponseStruct(t, "Blocks.Epoch", responseStruct, &BlocksEpoch{
		Success: true,
		Epoch:   "2017-03-21T13:00:00.000Z",
	})
}

// Get the transfer transaction fee.
func TestBlocksService_Fee(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getFee", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "fee": 10000000
			}`)
	})

	responseStruct, response, err := client.Blocks.Fee(context.Background())
	testGeneralError(t, "Blocks.Fee", err)
	testResponseUrl(t, "Blocks.Fee", response, "/api/blocks/getFee")
	testResponseStruct(t, "Blocks.Fee", responseStruct, &BlocksFee{
		Success: true,
		Fee:     10000000,
	})
}

// Get a list of transaction fees.
func TestBlocksService_Fees(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getFees", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "fees": {
			    "send": 10000000,
			    "vote": 100000000,
			    "secondsignature": 500000000,
			    "delegate": 2500000000,
			    "multisignature": 500000000
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Fees(context.Background())
	testGeneralError(t, "Blocks.Fees", err)
	testResponseUrl(t, "Blocks.Fees", response, "/api/blocks/getFees")
	testResponseStruct(t, "Blocks.Fees", responseStruct, &BlocksFees{
		Success: true,
		Fees: BlockFeeTypes{
			Send:            10000000,
			Vote:            100000000,
			Secondsignature: 500000000,
			Delegate:        2500000000,
			Multisignature:  500000000,
		},
	})
}

// Get the blockchain height.
func TestBlocksService_Height(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getHeight", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "height": 10
			}`)
	})

	responseStruct, response, err := client.Blocks.Height(context.Background())
	testGeneralError(t, "Blocks.Height", err)
	testResponseUrl(t, "Blocks.Height", response, "/api/blocks/getHeight")
	testResponseStruct(t, "Blocks.Height", responseStruct, &BlocksHeight{
		Success: true,
		Height:  10,
	})
}

// Get the blockchain milest
func TestBlocksService_Milestone(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getMilestone", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "milestone": 10
			}`)
	})

	responseStruct, response, err := client.Blocks.Milestone(context.Background())
	testGeneralError(t, "Blocks.Milestone", err)
	testResponseUrl(t, "Blocks.Milestone", response, "/api/blocks/getMilestone")
	testResponseStruct(t, "Blocks.Milestone", responseStruct, &BlocksMilestone{
		Success:   true,
		Milestone: 10,
	})
}

// Get the blockchain nethash.
func TestBlocksService_Nethash(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getNethash", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "nethash": "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23"
			}`)
	})

	responseStruct, response, err := client.Blocks.Nethash(context.Background())
	testGeneralError(t, "Blocks.Nethash", err)
	testResponseUrl(t, "Blocks.Nethash", response, "/api/blocks/getNethash")
	testResponseStruct(t, "Blocks.Nethash", responseStruct, &BlocksNethash{
		Success: true,
		Nethash: "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23",
	})
}

// Get the blockchain reward.
func TestBlocksService_Reward(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getReward", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "reward": 200000000
			}`)
	})

	responseStruct, response, err := client.Blocks.Reward(context.Background())
	testGeneralError(t, "Blocks.Reward", err)
	testResponseUrl(t, "Blocks.Reward", response, "/api/blocks/getReward")
	testResponseStruct(t, "Blocks.Reward", responseStruct, &BlocksReward{
		Success: true,
		Reward:  200000000,
	})
}

// Get the blockchain status.
func TestBlocksService_Status(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getStatus", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "epoch": "2017-03-21T13:00:00.000Z",
			  "height": 10,
			  "fee": 10000000,
			  "milestone": 10,
			  "nethash": "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23",
			  "reward": 200000000,
			  "supply": 12978931000000000
			}`)
	})

	responseStruct, response, err := client.Blocks.Status(context.Background())
	testGeneralError(t, "Blocks.Status", err)
	testResponseUrl(t, "Blocks.Status", response, "/api/blocks/getStatus")
	testResponseStruct(t, "Blocks.Status", responseStruct, &BlocksStatus{
		Success:   true,
		Epoch:     "2017-03-21T13:00:00.000Z",
		Height:    10,
		Fee:       10000000,
		Milestone: 10,
		Nethash:   "578e820911f24e039733b45e4882b73e301f813a0d2c31330dafda84534ffa23",
		Reward:    200000000,
		Supply:    12978931000000000,
	})
}

// Get the blockchain supply.
func TestBlocksService_Supply(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/getSupply", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "supply": 12978929800000000
			}`)
	})

	responseStruct, response, err := client.Blocks.Supply(context.Background())
	testGeneralError(t, "Blocks.Supply", err)
	testResponseUrl(t, "Blocks.Supply", response, "/api/blocks/getSupply")
	testResponseStruct(t, "Blocks.Supply", responseStruct, &BlocksSupply{
		Success: true,
		Supply:  12978929800000000,
	})
}

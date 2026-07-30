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

// Get all blocks.
func TestBlocksService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/blocks?page=1&limit=1",
			    "first": "/api/blocks?page=1&limit=1",
			    "last": "/api/blocks?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			      "version": 1,
			      "number": 1204291,
			      "parentHash": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			      "reward": "200000000",
			      "fee": "0",
			      "total": "200000000",
			      "amount": "0",
			      "transactionsRoot": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			      "payloadSize": 0,
			      "proposer": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			      "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			      "username": "thamar",
			      "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			      "confirmations": 55,
			      "transactionsCount": 0,
			      "timestamp": "1719434741918"
			    }
			  ]
			}`)
	})

	query := &BlocksQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Blocks.List(context.Background(), query)
	testGeneralError(t, "Blocks.List", err)
	testResponseUrl(t, "Blocks.List", response, "/api/blocks")
	testResponseStruct(t, "Blocks.List", responseStruct, &Blocks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/blocks?page=1&limit=1",
			First:      "/api/blocks?page=1&limit=1",
			Last:       "/api/blocks?page=1&limit=1",
		},
		Data: []Block{{
			Hash:              "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:           1,
			Number:            1204291,
			ParentHash:        "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Reward:            newBigInt(200000000),
			Fee:               newBigInt(0),
			Total:             newBigInt(200000000),
			Amount:            newBigInt(0),
			TransactionsRoot:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			PayloadSize:       0,
			Proposer:          "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			PublicKey:         "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			Username:          "thamar",
			Signature:         "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations:     55,
			TransactionsCount: 0,
			Timestamp:         "1719434741918",
		}},
	})
}

// Get a block by the given hash.
func TestBlocksService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "confirmations": 55,
			    "amount": "0",
			    "fee": "0",
			    "reward": "200000000",
			    "total": "200000000",
			    "proposer": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			    "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			    "username": "thamar",
			    "number": 1204291,
			    "hash": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			    "transactionsRoot": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			    "payloadSize": 0,
			    "parentHash": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			    "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			    "timestamp": "1719434741918",
			    "transactionsCount": 0,
			    "version": 1
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Get(context.Background(), "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96")
	testGeneralError(t, "Blocks.Get", err)
	testResponseUrl(t, "Blocks.Get", response, "/blocks/2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96")
	testResponseStruct(t, "Blocks.Get", responseStruct, &GetBlock{
		Data: Block{
			Hash:              "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:           1,
			Number:            1204291,
			ParentHash:        "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Reward:            newBigInt(200000000),
			Fee:               newBigInt(0),
			Total:             newBigInt(200000000),
			Amount:            newBigInt(0),
			TransactionsRoot:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			PayloadSize:       0,
			Proposer:          "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			PublicKey:         "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			Username:          "thamar",
			Signature:         "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations:     55,
			TransactionsCount: 0,
			Timestamp:         "1719434741918",
		},
	})
}

// Get the first block.
func TestBlocksService_First(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/first", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "confirmations": 4347586,
			    "amount": "12500000000000004",
			    "fee": "0",
			    "reward": "0",
			    "total": "0",
			    "proposer": "D6Z26L69gdk9qYmTv5uzk3uGepigtHY4ax",
			    "publicKey": "03d3fdad9c5b25bf8880e6b519eb3611a5c0b31adebc8455f0e096175b28321aff",
			    "number": 1,
			    "hash": "13114381566690093367",
			    "transactionsRoot": "2a44f340d76ffc3df204c5f38cd355b7496c9065a1ade2ef92071436bd72e867",
			    "payloadSize": 11395,
			    "parentHash": "0",
			    "signature": "3044022035694a9b99a9236655c658eb07fc3b02ce5edcc24b76424a7287c54ed3822b0602203621e92defb360490610f763d85e94c2db2807a4bd7756cc8a6a585463ef7bae",
			    "timestamp": "1490101200",
			    "transactionsCount": 52,
			    "version": 0
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.First(context.Background())
	testGeneralError(t, "Blocks.First", err)
	testResponseUrl(t, "Blocks.First", response, "/blocks/first")
	testResponseStruct(t, "Blocks.First", responseStruct, &GetBlock{
		Data: Block{
			Hash:              "13114381566690093367",
			Version:           0,
			Number:            1,
			ParentHash:        "0",
			Reward:            newBigInt(0),
			Fee:               newBigInt(0),
			Total:             newBigInt(0),
			Amount:            newBigInt(12500000000000004),
			TransactionsRoot:  "2a44f340d76ffc3df204c5f38cd355b7496c9065a1ade2ef92071436bd72e867",
			PayloadSize:       11395,
			Proposer:          "D6Z26L69gdk9qYmTv5uzk3uGepigtHY4ax",
			PublicKey:         "03d3fdad9c5b25bf8880e6b519eb3611a5c0b31adebc8455f0e096175b28321aff",
			Signature:         "3044022035694a9b99a9236655c658eb07fc3b02ce5edcc24b76424a7287c54ed3822b0602203621e92defb360490610f763d85e94c2db2807a4bd7756cc8a6a585463ef7bae",
			Confirmations:     4347586,
			TransactionsCount: 52,
			Timestamp:         "1490101200",
		},
	})
}

// Get the last block.
func TestBlocksService_Last(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/last", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "confirmations": 55,
			    "amount": "0",
			    "fee": "0",
			    "reward": "200000000",
			    "total": "200000000",
			    "proposer": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			    "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			    "username": "thamar",
			    "number": 1204291,
			    "hash": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			    "transactionsRoot": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			    "payloadSize": 0,
			    "parentHash": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			    "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			    "timestamp": "1719434741918",
			    "transactionsCount": 0,
			    "version": 1
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Last(context.Background())
	testGeneralError(t, "Blocks.Last", err)
	testResponseUrl(t, "Blocks.Last", response, "/blocks/last")
	testResponseStruct(t, "Blocks.Last", responseStruct, &GetBlock{
		Data: Block{
			Hash:              "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:           1,
			Number:            1204291,
			ParentHash:        "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Reward:            newBigInt(200000000),
			Fee:               newBigInt(0),
			Total:             newBigInt(200000000),
			Amount:            newBigInt(0),
			TransactionsRoot:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			PayloadSize:       0,
			Proposer:          "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			PublicKey:         "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			Username:          "thamar",
			Signature:         "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations:     55,
			TransactionsCount: 0,
			Timestamp:         "1719434741918",
		},
	})
}

// Get all transactions by the given block.
func TestBlocksService_Transactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1",
			    "first": "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1",
			    "last": "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1"
			  },
			  "data": [
			    {
					"hash": "b375b16677617ae903372040e6e794e239a043ac2017e2f944886cc7aaa9f3e9",
					"blockHash": "478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95",
					"value": "10000",
					"gas": "21000",
					"gasPrice": "10000000",
					"senderPublicKey": "03d69da6d6a2df142bf3af248881a76a01adaefd9a900d5f0116efdee763064876",
					"to": "DArvWfH5nMDT38tWmo5k461vMQpRXHQWX9",
					"from": "DDApe7WsjaAmvyPQhTeMWNq8iTv3JYD3zA",
					"data": "0x",
					"signature": "2c0468b15e86363c497acb5b62eb8e1274be6203af75eec5513ff906aa5e790bd6937ea9d37433cf50f345c4b9236cee5824e1ed6e0a62d829572fabf84d444e",
					"confirmations": 4858,
					"timestamp": "1719391504872",
					"nonce": "4",
					"receipt": {
						"cumulativeGasUsed": 21000,
						"gasRefunded": 0,
						"gasUsed": 21000,
						"status": 1
					}
			    }
			  ]
			}`)
	})

	query := &BlockTransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Blocks.Transactions(context.Background(), "478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95", query)
	testGeneralError(t, "Blocks.Transactions", err)
	testResponseUrl(t, "Blocks.Transactions", response, "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions")
	testResponseStruct(t, "Blocks.Transactions", responseStruct, &GetBlockTransactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1",
			First:      "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1",
			Last:       "/api/blocks/478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "b375b16677617ae903372040e6e794e239a043ac2017e2f944886cc7aaa9f3e9",
			BlockHash:       "478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95",
			Value:           newBigInt(10000),
			Gas:             newBigInt(21000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "03d69da6d6a2df142bf3af248881a76a01adaefd9a900d5f0116efdee763064876",
			To:              "DArvWfH5nMDT38tWmo5k461vMQpRXHQWX9",
			From:            "DDApe7WsjaAmvyPQhTeMWNq8iTv3JYD3zA",
			Data:            "0x",
			Signature:       "2c0468b15e86363c497acb5b62eb8e1274be6203af75eec5513ff906aa5e790bd6937ea9d37433cf50f345c4b9236cee5824e1ed6e0a62d829572fabf84d444e",
			Confirmations:   4858,
			Timestamp:       "1719391504872",
			Nonce:           newBigInt(4),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		}},
	})
}

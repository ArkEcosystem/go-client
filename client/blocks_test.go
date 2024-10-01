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
			      "id": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			      "version": 1,
			      "height": 1204291,
			      "previous": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			      "forged": {
			        "reward": "200000000",
			        "fee": "0",
			        "total": "200000000",
			        "amount": "0"
			      },
			      "payload": {
			        "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			        "length": 0
			      },
			      "generator": {
			        "username": "thamar",
			        "address": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			        "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941"
			      },
			      "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			      "confirmations": 55,
			      "transactions": 0,
			      "timestamp": 1719434741918
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
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
			Id:       "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:  1,
			Height:   1204291,
			Previous: "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "thamar",
				Address:   "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
				PublicKey: "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			},
			Signature:     "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations: 55,
			Transactions:  0,
			Timestamp: 1719434741918,
		}},
	})
}

// Get a block by the given id (id and height are valid)
func TestBlocksService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/10", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "confirmations": 55,
			    "forged": {
			      "amount": "0",
			      "fee": "0",
			      "reward": "200000000",
			      "total": "200000000"
			    },
			    "generator": {
			      "address": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			      "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			      "username": "thamar"
			    },
			    "height": 1204291,
			    "id": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			    "payload": {
			      "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			      "length": 0
			    },
			    "previous": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			    "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			    "timestamp": 1719434741918,
			    "transactions": 0,
			    "version": 1
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Get(context.Background(), 10)
	testGeneralError(t, "Blocks.Get", err)
	testResponseUrl(t, "Blocks.Get", response, "/blocks/10")
	testResponseStruct(t, "Blocks.Get", responseStruct, &GetBlock{
		Data: Block{
			Id:       "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:  1,
			Height:   1204291,
			Previous: "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "thamar",
				Address:   "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
				PublicKey: "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			},
			Signature:     "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations: 55,
			Transactions:  0,
			Timestamp: 1719434741918,
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
			    "forged": {
			      "amount": "12500000000000004",
			      "fee": "0",
			      "reward": "0",
			      "total": "0"
			    },
			    "generator": {
			      "address": "D6Z26L69gdk9qYmTv5uzk3uGepigtHY4ax",
			      "publicKey": "03d3fdad9c5b25bf8880e6b519eb3611a5c0b31adebc8455f0e096175b28321aff"
			    },
			    "height": 1,
			    "id": "13114381566690093367",
			    "payload": {
			      "hash": "2a44f340d76ffc3df204c5f38cd355b7496c9065a1ade2ef92071436bd72e867",
			      "length": 11395
			    },
			    "previous": "0",
			    "signature": "3044022035694a9b99a9236655c658eb07fc3b02ce5edcc24b76424a7287c54ed3822b0602203621e92defb360490610f763d85e94c2db2807a4bd7756cc8a6a585463ef7bae",
			    "timestamp": 1490101200,
			    "transactions": 52,
			    "version": 0
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.First(context.Background())
	testGeneralError(t, "Blocks.First", err)
	testResponseUrl(t, "Blocks.First", response, "/blocks/first")
	testResponseStruct(t, "Blocks.First", responseStruct, &GetBlock{
		Data: Block{
			Id:       "13114381566690093367",
			Version:  0,
			Height:   1,
			Previous: "0",
			Forged: BlockForged{
				Reward: 0,
				Fee:    0,
				Total:  0,
				Amount: 12500000000000004,
			},
			Payload: BlockPayload{
				Hash:   "2a44f340d76ffc3df204c5f38cd355b7496c9065a1ade2ef92071436bd72e867",
				Length: 11395,
			},
			Generator: BlockGenerator{
				Address:   "D6Z26L69gdk9qYmTv5uzk3uGepigtHY4ax",
				PublicKey: "03d3fdad9c5b25bf8880e6b519eb3611a5c0b31adebc8455f0e096175b28321aff",
			},
			Signature:    "3044022035694a9b99a9236655c658eb07fc3b02ce5edcc24b76424a7287c54ed3822b0602203621e92defb360490610f763d85e94c2db2807a4bd7756cc8a6a585463ef7bae",
			Confirmations: 4347586,
			Transactions: 52,
			Timestamp: 1490101200,
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
			    "forged": {
			      "amount": "0",
			      "fee": "0",
			      "reward": "200000000",
			      "total": "200000000"
			    },
			    "generator": {
			      "address": "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
			      "publicKey": "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			      "username": "thamar"
			    },
			    "height": 1204291,
			    "id": "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			    "payload": {
			      "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			      "length": 0
			    },
			    "previous": "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			    "signature": "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			    "timestamp": 1719434741918,
			    "transactions": 0,
			    "version": 1
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Last(context.Background())
	testGeneralError(t, "Blocks.Last", err)
	testResponseUrl(t, "Blocks.Last", response, "/blocks/last")
	testResponseStruct(t, "Blocks.Last", responseStruct, &GetBlock{
		Data: Block{
			Id:       "2bfefe91649a3df2a122f130d0ba6abf2be742db78bda5f590c753c53999ba96",
			Version:  1,
			Height:   1204291,
			Previous: "37e3bd18963762b0afbb35546daf4810feed9eebca8fe2348f4234d638ad756a",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "thamar",
				Address:   "DJGpX9UWf6v9goB73Z7CLL4gN6Pm54CoVm",
				PublicKey: "029283adfe1a02bf652c987180db5925d8c095c6ac5c4194fa844ef6a1f2ced941",
			},
			Signature:     "8084e6120d620fa1cd30068f74dc7e4a466670bea710efb95d9a109967e26510e3c87c587e590049dff2ce6b69b19468032abef8a60a2bfafc9696338abd5160b4a7e1839b969087d7dc346f7739bdc41e49404402d6db57839837d1a3241b70",
			Confirmations: 55,
			Transactions:  0,
			Timestamp: 1719434741918,
		},
	})
}

// Get all transactions by the given block.
func TestBlocksService_Transactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/10/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/blocks/10/transactions?page=1&limit=1",
			    "first": "/api/blocks/10/transactions?page=1&limit=1",
			    "last": "/api/blocks/10/transactions?page=1&limit=1"
			  },
			  "data": [
			    {
					"id": "b375b16677617ae903372040e6e794e239a043ac2017e2f944886cc7aaa9f3e9",
					"blockId": "478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95",
					"version": 1,
					"type": 0,
					"typeGroup": 1,
					"amount": "10000",
					"fee": "10000000",
					"senderPublicKey": "03d69da6d6a2df142bf3af248881a76a01adaefd9a900d5f0116efdee763064876",
					"recipient": "DArvWfH5nMDT38tWmo5k461vMQpRXHQWX9",
					"signature": "2c0468b15e86363c497acb5b62eb8e1274be6203af75eec5513ff906aa5e790bd6937ea9d37433cf50f345c4b9236cee5824e1ed6e0a62d829572fabf84d444e",
					"confirmations": 4858,
					"timestamp": 1719391504872,
					"nonce": "4",
					"signatures": null,
					"vendorField": null
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Blocks.Transactions(context.Background(), 10, query)
	testGeneralError(t, "Blocks.Transactions", err)
	testResponseUrl(t, "Blocks.Transactions", response, "/api/blocks/10/transactions")
	testResponseStruct(t, "Blocks.Transactions", responseStruct, &GetBlockTransactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/blocks/10/transactions?page=1&limit=1",
			First:      "/api/blocks/10/transactions?page=1&limit=1",
			Last:       "/api/blocks/10/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:              "b375b16677617ae903372040e6e794e239a043ac2017e2f944886cc7aaa9f3e9",
			BlockId:         "478ccd06f1f93b74e6ec5fc626ffe8d3c2421b3e074d070caa3c0ab4e8a6ad95",
			Version:         1,
			Type:            0,
			TypeGroup:       1,
			Amount:          10000,
			Fee:             10000000,
			SenderPublicKey: "03d69da6d6a2df142bf3af248881a76a01adaefd9a900d5f0116efdee763064876",
			Recipient:       "DArvWfH5nMDT38tWmo5k461vMQpRXHQWX9",
			Signature:       "2c0468b15e86363c497acb5b62eb8e1274be6203af75eec5513ff906aa5e790bd6937ea9d37433cf50f345c4b9236cee5824e1ed6e0a62d829572fabf84d444e",
			Confirmations:   4858,
			Timestamp:       1719391504872,
			Nonce:           4,
			Signatures:      nil,
			VendorField:     "",
		}},
	})
}

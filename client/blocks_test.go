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
			      "id": "dummy",
			      "version": 0,
			      "height": 10,
			      "previous": "dummy",
			      "forged": {
			        "reward": "200000000",
			        "fee": "0",
			        "total": "200000000",
			        "amount": "0"
			      },
			      "payload": {
			        "hash": "dummy",
			        "length": 0
			      },
			      "generator": {
			        "username": "dummy",
			        "address": "dummy",
			        "publicKey": "dummy"
			      },
			      "signature": "dummy",
			      "confirmations": 0,
			      "transactions": 0,
			      "timestamp": {
			        "epoch": 40678848,
			        "unix": 1530780048,
			        "human": "2018-07-05T08:40:48Z"
			      }
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
			Next:       "",
			Previous:   "",
			Self:       "/api/blocks?page=1&limit=1",
			First:      "/api/blocks?page=1&limit=1",
			Last:       "/api/blocks?page=1&limit=1",
		},
		Data: []Block{{
			Id:       "dummy",
			Version:  0,
			Height:   10,
			Previous: "dummy",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "dummy",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "dummy",
				Address:   "dummy",
				PublicKey: "dummy",
			},
			Signature:     "dummy",
			Confirmations: 0,
			Transactions:  0,
			Timestamp: Timestamp{
				Epoch: 40678848,
				Unix:  1530780048,
				Human: "2018-07-05T08:40:48Z",
			},
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
			    "id": "dummy",
			    "version": 0,
			    "height": 10,
			    "previous": "dummy",
			    "forged": {
			      "reward": "200000000",
			      "fee": "0",
			      "total": "200000000",
			      "amount": "0"
			    },
			    "payload": {
			      "hash": "dummy",
			      "length": 0
			    },
			    "generator": {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy"
			    },
			    "signature": "dummy",
			    "confirmations": 0,
			    "transactions": 0,
			    "timestamp": {
			      "epoch": 40678848,
			      "unix": 1530780048,
			      "human": "2018-07-05T08:40:48Z"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Get(context.Background(), 10)
	testGeneralError(t, "Blocks.Get", err)
	testResponseUrl(t, "Blocks.Get", response, "/blocks/10")
	testResponseStruct(t, "Blocks.Get", responseStruct, &GetBlock{
		Data: Block{
			Id:       "dummy",
			Version:  0,
			Height:   10,
			Previous: "dummy",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "dummy",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "dummy",
				Address:   "dummy",
				PublicKey: "dummy",
			},
			Signature:     "dummy",
			Confirmations: 0,
			Transactions:  0,
			Timestamp: Timestamp{
				Epoch: 40678848,
				Unix:  1530780048,
				Human: "2018-07-05T08:40:48Z",
			},
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
			    "id": "13114381566690093367",
			    "version": 0,
			    "height": 1,
			    "previous": "0",
			    "forged": {
			      "reward": "0",
			      "fee": "0",
			      "total": "0",
			      "amount": "12500000000000004"
			    },
			    "payload": {
			      "hash": "2a44f340d76ffc3df204c5f38cd355b7496c9065a1ade2ef92071436bd72e867",
			      "length": 11395
			    },
			    "generator": {
			      "address": "D6Z26L69gdk9qYmTv5uzk3uGepigtHY4ax",
			      "publicKey": "03d3fdad9c5b25bf8880e6b519eb3611a5c0b31adebc8455f0e096175b28321aff"
			    },
			    "signature": "3044022035694a9b99a9236655c658eb07fc3b02ce5edcc24b76424a7287c54ed3822b0602203621e92defb360490610f763d85e94c2db2807a4bd7756cc8a6a585463ef7bae",
			    "confirmations": 4347586,
			    "transactions": 52,
			    "timestamp": {
			      "epoch": 0,
			      "unix": 1490101200,
			      "human": "2017-03-21T13:00:00.000Z"
			    }
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
			Timestamp: Timestamp{
				Epoch: 0,
				Unix:  1490101200,
				Human: "2017-03-21T13:00:00.000Z",
			},
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
			    "id": "dummy",
			    "version": 0,
			    "height": 10,
			    "previous": "dummy",
			    "forged": {
			      "reward": "200000000",
			      "fee": "0",
			      "total": "200000000",
			      "amount": "0"
			    },
			    "payload": {
			      "hash": "dummy",
			      "length": 0
			    },
			    "generator": {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy"
			    },
			    "signature": "dummy",
			    "confirmations": 0,
			    "transactions": 0,
			    "timestamp": {
			      "epoch": 40678848,
			      "unix": 1530780048,
			      "human": "2018-07-05T08:40:48Z"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Blocks.Last(context.Background())
	testGeneralError(t, "Blocks.Last", err)
	testResponseUrl(t, "Blocks.Last", response, "/blocks/last")
	testResponseStruct(t, "Blocks.Last", responseStruct, &GetBlock{
		Data: Block{
			Id:       "dummy",
			Version:  0,
			Height:   10,
			Previous: "dummy",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "dummy",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "dummy",
				Address:   "dummy",
				PublicKey: "dummy",
			},
			Signature:     "dummy",
			Confirmations: 0,
			Transactions:  0,
			Timestamp: Timestamp{
				Epoch: 40678848,
				Unix:  1530780048,
				Human: "2018-07-05T08:40:48Z",
			},
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
			      "id": "dummy",
			      "blockId": "dummy",
			      "type": 0,
			      "typeGroup": 1,
			      "amount": "10000000",
			      "fee": "10000000",
			      "sender": "dummy",
			      "senderPublicKey": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10348,
			      "timestamp": {
			        "epoch": 40505460,
			        "unix": 1530606660,
			        "human": "2018-07-03T08:31:00Z"
			      },
			      "nonce": "1"
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
			Next:       "",
			Previous:   "",
			Self:       "/api/blocks/10/transactions?page=1&limit=1",
			First:      "/api/blocks/10/transactions?page=1&limit=1",
			Last:       "/api/blocks/10/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:              "dummy",
			BlockId:         "dummy",
			Type:            0,
			TypeGroup:       1,
			Amount:          10000000,
			Fee:             10000000,
			Sender:          "dummy",
			SenderPublicKey: "dummy",
			Recipient:       "dummy",
			Signature:       "dummy",
			VendorField:     "dummy",
			Confirmations:   10348,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
			Nonce: 1,
		}},
	})
}

// Filter all blocks by the given criteria.
func TestBlocksService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blocks/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/blocks/search?page=1&limit=1",
			    "first": "/api/blocks/search?page=1&limit=1",
			    "last": "/api/blocks/search?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummy",
			      "version": 0,
			      "height": 10,
			      "previous": "dummy",
			      "forged": {
			        "reward": "200000000",
			        "fee": "0",
			        "total": "200000000",
			        "amount": "0"
			      },
			      "payload": {
			        "hash": "dummy",
			        "length": 0
			      },
			      "generator": {
			        "username": "dummy",
			        "address": "dummy",
			        "publicKey": "dummy"
			      },
			      "signature": "dummy",
			      "confirmations": 0,
			      "transactions": 0,
			      "timestamp": {
			        "epoch": 40678848,
			        "unix": 1530780048,
			        "human": "2018-07-05T08:40:48Z"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &BlocksSearchRequest{Id: "dummy"}
	responseStruct, response, err := client.Blocks.Search(context.Background(), query, body)
	testGeneralError(t, "Blocks.Search", err)
	testResponseUrl(t, "Blocks.Search", response, "/api/blocks/search")
	testResponseStruct(t, "Blocks.Search", responseStruct, &Blocks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/blocks/search?page=1&limit=1",
			First:      "/api/blocks/search?page=1&limit=1",
			Last:       "/api/blocks/search?page=1&limit=1",
		},
		Data: []Block{{
			Id:       "dummy",
			Version:  0,
			Height:   10,
			Previous: "dummy",
			Forged: BlockForged{
				Reward: 200000000,
				Fee:    0,
				Total:  200000000,
				Amount: 0,
			},
			Payload: BlockPayload{
				Hash:   "dummy",
				Length: 0,
			},
			Generator: BlockGenerator{
				Username:  "dummy",
				Address:   "dummy",
				PublicKey: "dummy",
			},
			Signature:     "dummy",
			Confirmations: 0,
			Transactions:  0,
			Timestamp: Timestamp{
				Epoch: 40678848,
				Unix:  1530780048,
				Human: "2018-07-05T08:40:48Z",
			},
		}},
	})
}

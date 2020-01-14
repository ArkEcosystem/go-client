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

// Get all accounts.
func TestDelegatesService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/delegates?page=1&limit=1",
			    "first": "/api/delegates?page=1&limit=1",
			    "last": "/api/delegates?page=1&limit=1"
			  },
			  "data": [
			    {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy",
			      "votes": 1000,
			      "rank": 1,
			      "blocks": {
			        "produced": 119,
			        "missed": 56,
			        "last": {
			          "id": "dummy",
			          "timestamp": {
			            "epoch": 40686944,
			            "unix": 1530788144,
			            "human": "2018-07-05T10:55:44Z"
			          }
			        }
			      },
			      "production": {
			        "approval": 0.01
			      },
			      "forged": {
			        "fees": "468407250508",
			        "rewards": "13589400000000",
			        "total": "14057807250508"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Delegates.List(context.Background(), query)
	testGeneralError(t, "Delegates.List", err)
	testResponseUrl(t, "Delegates.List", response, "/api/delegates")
	testResponseStruct(t, "Delegates.List", responseStruct, &Delegates{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/delegates?page=1&limit=1",
			First:      "/api/delegates?page=1&limit=1",
			Last:       "/api/delegates?page=1&limit=1",
		},
		Data: []Delegate{{
			Username:  "dummy",
			Address:   "dummy",
			PublicKey: "dummy",
			Votes:     1000,
			Rank:      1,
			Blocks: DelegateBlocks{
				Produced: 119,
				Missed:   56,
				Last: Block{
					Id: "dummy",
					Timestamp: Timestamp{
						Epoch: 40686944,
						Unix:  1530788144,
						Human: "2018-07-05T10:55:44Z",
					},
				},
			},
			Production: DelegateProduction{
				Approval: 0.01,
			},
			Forged: DelegateForged{
				Fees:    468407250508,
				Rewards: 13589400000000,
				Total:   14057807250508,
			},
		}},
	})
}

// Get a delegate by the given ID (address, publicKey and username are valid)
func TestDelegatesService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/delegates/dummy?page=1&limit=1",
			    "first": "/api/delegates/dummy?page=1&limit=1",
			    "last": "/api/delegates/dummy?page=1&limit=1"
			  },
			  "data": {
			    "username": "dummy",
			    "address": "dummy",
			    "publicKey": "dummy",
			    "votes": 1000,
			    "rank": 1,
			    "blocks": {
			      "produced": 119,
			      "missed": 56,
			      "last": {
			        "id": "dummy",
			        "timestamp": {
			          "epoch": 40686944,
			          "unix": 1530788144,
			          "human": "2018-07-05T10:55:44Z"
			        }
			      }
			    },
			    "production": {
			      "approval": 0.01
			    },
			    "forged": {
			      "fees": "468407250508",
			      "rewards": "13589400000000",
			      "total": "14057807250508"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Delegates.Get(context.Background(), "dummy")
	testGeneralError(t, "Delegates.Get", err)
	testResponseUrl(t, "Delegates.Get", response, "/api/delegates/dummy")
	testResponseStruct(t, "Delegates.Get", responseStruct, &GetDelegate{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/delegates/dummy?page=1&limit=1",
			First:      "/api/delegates/dummy?page=1&limit=1",
			Last:       "/api/delegates/dummy?page=1&limit=1",
		},
		Data: Delegate{
			Username:  "dummy",
			Address:   "dummy",
			PublicKey: "dummy",
			Votes:     1000,
			Rank:      1,
			Blocks: DelegateBlocks{
				Produced: 119,
				Missed:   56,
				Last: Block{
					Id: "dummy",
					Timestamp: Timestamp{
						Epoch: 40686944,
						Unix:  1530788144,
						Human: "2018-07-05T10:55:44Z",
					},
				},
			},
			Production: DelegateProduction{
				Approval: 0.01,
			},
			Forged: DelegateForged{
				Fees:    468407250508,
				Rewards: 13589400000000,
				Total:   14057807250508,
			},
		},
	})
}

// Get all blocks for the given delegate.
func TestDelegatesService_Blocks(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/dummy/blocks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/delegates/dummy/blocks?page=1&limit=1",
			    "first": "/api/delegates/dummy/blocks?page=1&limit=1",
			    "last": "/api/delegates/dummy/blocks?page=1&limit=1"
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
	responseStruct, response, err := client.Delegates.Blocks(context.Background(), "dummy", query)
	testGeneralError(t, "Delegates.Blocks", err)
	testResponseUrl(t, "Delegates.Blocks", response, "/api/delegates/dummy/blocks")
	testResponseStruct(t, "Delegates.Blocks", responseStruct, &GetDelegateBlocks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/delegates/dummy/blocks?page=1&limit=1",
			First:      "/api/delegates/dummy/blocks?page=1&limit=1",
			Last:       "/api/delegates/dummy/blocks?page=1&limit=1",
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
			Signature:    "dummy",
			Transactions: 0,
			Timestamp: Timestamp{
				Epoch: 40678848,
				Unix:  1530780048,
				Human: "2018-07-05T08:40:48Z",
			},
		}},
	})
}

// Get all voters for the given delegate.
func TestDelegatesService_Voters(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/dummy/voters", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/delegates/dummy/blocks?page=1&limit=1",
			    "first": "/api/delegates/dummy/blocks?page=1&limit=1",
			    "last": "/api/delegates/dummy/blocks?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "nonce": "1",
			      "balance": "100000000",
			      "isDelegate": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Delegates.Voters(context.Background(), "dummy", query)
	testGeneralError(t, "Delegates.Voters", err)
	testResponseUrl(t, "Delegates.Voters", response, "/api/delegates/dummy/voters")
	testResponseStruct(t, "Delegates.Voters", responseStruct, &GetDelegateVoters{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/delegates/dummy/blocks?page=1&limit=1",
			First:      "/api/delegates/dummy/blocks?page=1&limit=1",
			Last:       "/api/delegates/dummy/blocks?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:    "dummy",
			PublicKey:  "dummy",
			Nonce:      1,
			Balance:    100000000,
			IsDelegate: false,
		}},
	})
}

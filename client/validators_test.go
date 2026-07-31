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
func TestValidatorsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/validators", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/validators?page=1&limit=1",
			    "first": "/api/validators?page=1&limit=1",
			    "last": "/api/validators?page=1&limit=1"
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
			          "timestamp": 1719434741918
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
	responseStruct, response, err := client.Validators.List(context.Background(), query)
	testGeneralError(t, "Validators.List", err)
	testResponseUrl(t, "Validators.List", response, "/api/validators")
	testResponseStruct(t, "Validators.List", responseStruct, &Validators{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/validators?page=1&limit=1",
			First:      "/api/validators?page=1&limit=1",
			Last:       "/api/validators?page=1&limit=1",
		},
		Data: []Validator{{
			Username:  "dummy",
			Address:   "dummy",
			PublicKey: "dummy",
			Votes:     1000,
			Rank:      1,
			Blocks: ValidatorBlocks{
				Produced: 119,
				Missed:   56,
				Last: Block{
					Id:        "dummy",
					Timestamp: 1719434741918,
				},
			},
			Production: ValidatorProduction{
				Approval: 0.01,
			},
			Forged: ValidatorForged{
				Fees:    468407250508,
				Rewards: 13589400000000,
				Total:   14057807250508,
			},
		}},
	})
}

// Get a validator by the given ID (address, publicKey and username are valid)
func TestValidatorsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/validators/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/validators/dummy?page=1&limit=1",
			    "first": "/api/validators/dummy?page=1&limit=1",
			    "last": "/api/validators/dummy?page=1&limit=1"
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
			        "timestamp": 1719434741918
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

	responseStruct, response, err := client.Validators.Get(context.Background(), "dummy")
	testGeneralError(t, "Validators.Get", err)
	testResponseUrl(t, "Validators.Get", response, "/api/validators/dummy")
	testResponseStruct(t, "Validators.Get", responseStruct, &GetValidator{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/validators/dummy?page=1&limit=1",
			First:      "/api/validators/dummy?page=1&limit=1",
			Last:       "/api/validators/dummy?page=1&limit=1",
		},
		Data: Validator{
			Username:  "dummy",
			Address:   "dummy",
			PublicKey: "dummy",
			Votes:     1000,
			Rank:      1,
			Blocks: ValidatorBlocks{
				Produced: 119,
				Missed:   56,
				Last: Block{
					Id:        "dummy",
					Timestamp: 1719434741918,
				},
			},
			Production: ValidatorProduction{
				Approval: 0.01,
			},
			Forged: ValidatorForged{
				Fees:    468407250508,
				Rewards: 13589400000000,
				Total:   14057807250508,
			},
		},
	})
}

// Get all blocks for the given validator.
func TestValidatorsService_Blocks(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/validators/dummy/blocks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/validators/dummy/blocks?page=1&limit=1",
			    "first": "/api/validators/dummy/blocks?page=1&limit=1",
			    "last": "/api/validators/dummy/blocks?page=1&limit=1"
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
			      "timestamp": 1719434741918
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Validators.Blocks(context.Background(), "dummy", query)
	testGeneralError(t, "Validators.Blocks", err)
	testResponseUrl(t, "Validators.Blocks", response, "/api/validators/dummy/blocks")
	testResponseStruct(t, "Validators.Blocks", responseStruct, &GetValidatorBlocks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/validators/dummy/blocks?page=1&limit=1",
			First:      "/api/validators/dummy/blocks?page=1&limit=1",
			Last:       "/api/validators/dummy/blocks?page=1&limit=1",
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
			Timestamp:    1719434741918,
		}},
	})
}

// Get all voters for the given validator.
func TestValidatorsService_Voters(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/validators/dummy/voters", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/validators/dummy/voters?page=1&limit=1",
			    "first": "/api/validators/dummy/voters?page=1&limit=1",
			    "last": "/api/validators/dummy/voters?page=1&limit=1"
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
	responseStruct, response, err := client.Validators.Voters(context.Background(), "dummy", query)
	testGeneralError(t, "Validators.Voters", err)
	testResponseUrl(t, "Validators.Voters", response, "/api/validators/dummy/voters")
	testResponseStruct(t, "Validators.Voters", responseStruct, &GetValidatorVoters{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/validators/dummy/voters?page=1&limit=1",
			First:      "/api/validators/dummy/voters?page=1&limit=1",
			Last:       "/api/validators/dummy/voters?page=1&limit=1",
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

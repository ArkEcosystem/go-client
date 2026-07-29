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

// Get all validators.
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
			      "address": "dummy",
			      "publicKey": "dummy",
			      "balance": "1000000000",
			      "nonce": "1",
			      "updated_at": "2026-03-02T15:00:00.000Z",
			      "tokenCount": 0,
			      "attributes": {
			        "username": "dummy",
			        "validatorRank": 1,
			        "validatorApproval": 0.01,
			        "validatorResigned": false,
			        "validatorPublicKey": "dummy",
			        "validatorForgedFees": "468407250508",
			        "validatorForgedRewards": "13589400000000",
			        "validatorForgedTotal": "14057807250508",
			        "validatorVoteBalance": "156947252547993",
			        "validatorVotersCount": 42,
			        "validatorProducedBlocks": 119,
			        "validatorLastBlock": {
			          "id": "dummy",
			          "height": 1204291,
			          "timestamp": 1719434741918
			        }
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Validators.List(context.Background(), query)
	testGeneralError(t, "Validators.List", err)
	testResponseUrl(t, "Validators.List", response, "/api/validators")
	testResponseStruct(t, "Validators.List", responseStruct, &Wallets{
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
		Data: []Wallet{{
			Address:    "dummy",
			PublicKey:  "dummy",
			Balance:    newBigInt(1000000000),
			Nonce:      newBigInt(1),
			UpdatedAt:  "2026-03-02T15:00:00.000Z",
			TokenCount: 0,
			Attributes: WalletAttributes{
				Username:                "dummy",
				ValidatorRank:           1,
				ValidatorApproval:       0.01,
				ValidatorResigned:       false,
				ValidatorPublicKey:      "dummy",
				ValidatorForgedFees:     newBigInt(468407250508),
				ValidatorForgedRewards:  newBigInt(13589400000000),
				ValidatorForgedTotal:    newBigInt(14057807250508),
				ValidatorVoteBalance:    newBigInt(156947252547993),
				ValidatorVotersCount:    42,
				ValidatorProducedBlocks: 119,
				ValidatorLastBlock: &ValidatorLastBlock{
					Id:        "dummy",
					Height:    1204291,
					Timestamp: 1719434741918,
				},
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
			  "data": {
			    "address": "dummy",
			    "publicKey": "dummy",
			    "balance": "1000000000",
			    "nonce": "1",
			    "attributes": {
			      "username": "dummy",
			      "validatorRank": 1,
			      "validatorForgedFees": "468407250508",
			      "validatorForgedRewards": "13589400000000",
			      "validatorForgedTotal": "14057807250508"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Validators.Get(context.Background(), "dummy")
	testGeneralError(t, "Validators.Get", err)
	testResponseUrl(t, "Validators.Get", response, "/api/validators/dummy")
	testResponseStruct(t, "Validators.Get", responseStruct, &GetWallet{
		Data: Wallet{
			Address:   "dummy",
			PublicKey: "dummy",
			Balance:   newBigInt(1000000000),
			Nonce:     newBigInt(1),
			Attributes: WalletAttributes{
				Username:               "dummy",
				ValidatorRank:          1,
				ValidatorForgedFees:    newBigInt(468407250508),
				ValidatorForgedRewards: newBigInt(13589400000000),
				ValidatorForgedTotal:   newBigInt(14057807250508),
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
			      "hash": "dummy",
			      "version": 0,
			      "number": 10,
			      "parentHash": "dummy",
			      "reward": "200000000",
			      "fee": "0",
			      "total": "200000000",
			      "amount": "0",
			      "proposer": "dummy",
			      "publicKey": "dummy",
			      "username": "dummy",
			      "signature": "dummy",
			      "transactionsCount": 0,
			      "timestamp": "1719434741918"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Validators.Blocks(context.Background(), "dummy", query)
	testGeneralError(t, "Validators.Blocks", err)
	testResponseUrl(t, "Validators.Blocks", response, "/api/validators/dummy/blocks")
	testResponseStruct(t, "Validators.Blocks", responseStruct, &Blocks{
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
			Hash:              "dummy",
			Version:           0,
			Number:            10,
			ParentHash:        "dummy",
			Reward:            newBigInt(200000000),
			Fee:               newBigInt(0),
			Total:             newBigInt(200000000),
			Amount:            newBigInt(0),
			Proposer:          "dummy",
			PublicKey:         "dummy",
			Username:          "dummy",
			Signature:         "dummy",
			TransactionsCount: 0,
			Timestamp:         "1719434741918",
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
			      "balance": "100000000"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Validators.Voters(context.Background(), "dummy", query)
	testGeneralError(t, "Validators.Voters", err)
	testResponseUrl(t, "Validators.Voters", response, "/api/validators/dummy/voters")
	testResponseStruct(t, "Validators.Voters", responseStruct, &Wallets{
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
			Address:   "dummy",
			PublicKey: "dummy",
			Nonce:     newBigInt(1),
			Balance:   newBigInt(100000000),
		}},
	})
}

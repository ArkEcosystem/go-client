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
			      "address": "0xe5a97E663158dEaF3b65bBF88897b8359Dc19F81",
			      "publicKey": "027b320c5429334ecf846122492d12b898a756bf1347aa61f7bf1dcd706315a9fb",
			      "balance": "934028115287775973804882",
			      "nonce": "0",
			      "updated_at": "23191163",
			      "attributes": {
			        "vote": "0xe5a97E663158dEaF3b65bBF88897b8359Dc19F81",
			        "isLegacy": true,
			        "username": "genesis_31",
			        "legacyNonce": "7",
			        "validatorFee": "0",
			        "validatorRank": 1,
			        "validatorApproval": 0.0084,
			        "validatorResigned": false,
			        "validatorLastBlock": {
			          "hash": "303cbc580d9d4c167cc7b1ecbb525fffb137ee58611fbd1474741f32340f9eb5",
			          "number": 23191163,
			          "timestamp": 1785333484026
			        },
			        "validatorPublicKey": "91ff20e1aee92c4e6febc1f7e1e55355d182812536055afb6a1bab300387580707bc0536e9d994e84fe58be8513e2550",
			        "validatorForgedFees": "5983395973804882",
			        "validatorForgedTotal": "17160005983395973804882",
			        "validatorVoteBalance": "1445434454042994148293705",
			        "validatorVotersCount": 23,
			        "validatorForgedRewards": "17160000000000000000000",
			        "validatorProducedBlocks": 10449
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
			Address:   "0xe5a97E663158dEaF3b65bBF88897b8359Dc19F81",
			PublicKey: "027b320c5429334ecf846122492d12b898a756bf1347aa61f7bf1dcd706315a9fb",
			Balance:   newBigIntFromString("934028115287775973804882"),
			Nonce:     newBigInt(0),
			UpdatedAt: "23191163",
			Attributes: WalletAttributes{
				Vote:                    "0xe5a97E663158dEaF3b65bBF88897b8359Dc19F81",
				IsLegacy:                true,
				Username:                "genesis_31",
				LegacyNonce:             "7",
				ValidatorFee:            newBigInt(0),
				ValidatorRank:           1,
				ValidatorApproval:       0.0084,
				ValidatorResigned:       false,
				ValidatorPublicKey:      "91ff20e1aee92c4e6febc1f7e1e55355d182812536055afb6a1bab300387580707bc0536e9d994e84fe58be8513e2550",
				ValidatorForgedFees:     newBigInt(5983395973804882),
				ValidatorForgedTotal:    newBigIntFromString("17160005983395973804882"),
				ValidatorVoteBalance:    newBigIntFromString("1445434454042994148293705"),
				ValidatorVotersCount:    23,
				ValidatorForgedRewards:  newBigIntFromString("17160000000000000000000"),
				ValidatorProducedBlocks: 10449,
				ValidatorLastBlock: &ValidatorLastBlock{
					Hash:      "303cbc580d9d4c167cc7b1ecbb525fffb137ee58611fbd1474741f32340f9eb5",
					Number:    23191163,
					Timestamp: 1785333484026,
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
			      "validatorFee": "0",
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
				ValidatorFee:           newBigInt(0),
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
			      "round": 2,
			      "parentHash": "dummy",
			      "reward": "200000000",
			      "fee": "0",
			      "gasUsed": 0,
			      "total": "200000000",
			      "amount": "0",
			      "proposer": "dummy",
			      "publicKey": "dummy",
			      "username": "dummy",
			      "validatorSet": "5207266566312791",
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
			Round:             2,
			ParentHash:        "dummy",
			Reward:            newBigInt(200000000),
			Fee:               newBigInt(0),
			GasUsed:           0,
			Total:             newBigInt(200000000),
			Amount:            newBigInt(0),
			Proposer:          "dummy",
			PublicKey:         "dummy",
			Username:          "dummy",
			ValidatorSet:      "5207266566312791",
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

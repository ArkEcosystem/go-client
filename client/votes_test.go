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

// Get all votes.
func TestVotesService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/votes", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/votes?page=1&limit=1",
			    "first": "/api/votes?page=1&limit=1",
			    "last": "/api/votes?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "0",
			      "gas": "100000",
			      "gasPrice": "10000000",
			      "senderPublicKey": "dummy",
			      "to": "dummy",
			      "from": "dummy",
			      "data": "0x",
			      "signature": "dummy",
			      "confirmations": 10,
			      "timestamp": "1719434741918",
			      "nonce": "1",
			      "receipt": {
			        "cumulativeGasUsed": 100000,
			        "gasRefunded": 0,
			        "gasUsed": 100000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &VotesQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Votes.List(context.Background(), query)
	testGeneralError(t, "Votes.List", err)
	testResponseUrl(t, "Votes.List", response, "/api/votes")
	testResponseStruct(t, "Votes.List", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/votes?page=1&limit=1",
			First:      "/api/votes?page=1&limit=1",
			Last:       "/api/votes?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(0),
			Gas:             newBigInt(100000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 100000,
				GasRefunded:       0,
				GasUsed:           100000,
				Status:            1,
			},
		}},
	})
}

// Get a vote by the given id.
func TestVotesService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/votes/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "hash": "dummy",
			    "blockHash": "dummy",
			    "value": "0",
			    "gas": "100000",
			    "gasPrice": "10000000",
			    "senderPublicKey": "dummy",
			    "to": "dummy",
			    "from": "dummy",
			    "data": "0x",
			    "signature": "dummy",
			    "confirmations": 10,
			    "timestamp": "1719434741918",
			    "nonce": "1",
			    "receipt": {
			      "cumulativeGasUsed": 100000,
			      "gasRefunded": 0,
			      "gasUsed": 100000,
			      "status": 1
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Votes.Get(context.Background(), "dummy", nil)
	testGeneralError(t, "Votes.Get", err)
	testResponseUrl(t, "Votes.Get", response, "/api/votes/dummy")
	testResponseStruct(t, "Votes.Get", responseStruct, &GetTransaction{
		Data: Transaction{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(0),
			Gas:             newBigInt(100000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 100000,
				GasRefunded:       0,
				GasUsed:           100000,
				Status:            1,
			},
		},
	})
}

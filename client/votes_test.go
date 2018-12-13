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
			      "id": "dummy",
			      "blockId": "dummy",
			      "type": 3,
			      "amount": 0,
			      "fee": 100000000,
			      "sender": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "asset": {
			        "votes": [
			          "+dummy"
			        ]
			      },
			      "confirmations": 10,
			      "timestamp": {
			        "epoch": 39862054,
			        "unix": 1529963254,
			        "human": "2018-06-25T21:47:34Z"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Votes.List(context.Background(), query)
	testGeneralError(t, "Votes.List", err)
	testResponseUrl(t, "Votes.List", response, "/api/votes")
	testResponseStruct(t, "Votes.List", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/votes?page=1&limit=1",
			First:      "/api/votes?page=1&limit=1",
			Last:       "/api/votes?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:        "dummy",
			BlockId:   "dummy",
			Type:      3,
			Amount:    0,
			Fee:       100000000,
			Sender:    "dummy",
			Recipient: "dummy",
			Signature: "dummy",
			Asset: &TransactionAsset{
				Votes: []string{
					"+dummy",
				},
			},
			Confirmations: 10,
			Timestamp: Timestamp{
				Epoch: 39862054,
				Unix:  1529963254,
				Human: "2018-06-25T21:47:34Z",
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
			    "id": "dummy",
			    "blockId": "dummy",
			    "type": 3,
			    "amount": 0,
			    "fee": 100000000,
			    "sender": "dummy",
			    "recipient": "dummy",
			    "signature": "dummy",
			    "asset": {
			      "votes": [
			        "+dummy"
			      ]
			    },
			    "confirmations": 10,
			    "timestamp": {
			      "epoch": 39862054,
			      "unix": 1529963254,
			      "human": "2018-06-25T21:47:34Z"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Votes.Get(context.Background(), "dummy")
	testGeneralError(t, "Votes.Get", err)
	testResponseUrl(t, "Votes.Get", response, "/api/votes/dummy")
	testResponseStruct(t, "Votes.Get", responseStruct, &GetTransaction{
		Data: Transaction{
			Id:        "dummy",
			BlockId:   "dummy",
			Type:      3,
			Amount:    0,
			Fee:       100000000,
			Sender:    "dummy",
			Recipient: "dummy",
			Signature: "dummy",
			Asset: &TransactionAsset{
				Votes: []string{
					"+dummy",
				},
			},
			Confirmations: 10,
			Timestamp: Timestamp{
				Epoch: 39862054,
				Unix:  1529963254,
				Human: "2018-06-25T21:47:34Z",
			},
		},
	})
}

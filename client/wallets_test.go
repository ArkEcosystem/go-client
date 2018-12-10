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

// Get all wallets.
func TestWalletsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets?page=1&limit=1",
			    "first": "/api/wallets?page=1&limit=1",
			    "last": "/api/wallets?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "balance": 1000000000,
			      "isDelegate": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.List(context.Background(), query)
	testGeneralError(t, "Wallets.List", err)
	testResponseUrl(t, "Wallets.List", response, "/api/wallets")
	testResponseStruct(t, "Wallets.List", responseStruct, &Wallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets?page=1&limit=1",
			First:      "/api/wallets?page=1&limit=1",
			Last:       "/api/wallets?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:    "dummy",
			PublicKey:  "dummy",
			Balance:    1000000000,
			IsDelegate: false,
		}},
	})
}

// Get all wallets sorted by balance in descending order.
func TestWalletsService_Top(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/top", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/top?page=1&limit=1",
			    "first": "/api/wallets/top?page=1&limit=1",
			    "last": "/api/wallets/top?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "balance": 1000000000,
			      "isDelegate": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.Top(context.Background(), query)
	testGeneralError(t, "Wallets.Top", err)
	testResponseUrl(t, "Wallets.Top", response, "/api/wallets/top")
	testResponseStruct(t, "Wallets.Top", responseStruct, &Wallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/top?page=1&limit=1",
			First:      "/api/wallets/top?page=1&limit=1",
			Last:       "/api/wallets/top?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:    "dummy",
			PublicKey:  "dummy",
			Balance:    1000000000,
			IsDelegate: false,
		}},
	})
}

// Get a wallet by the given id.
func TestWalletsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "address": "dummy",
			    "publicKey": "dummy",
			    "balance": 1000000000,
			    "isDelegate": false
			  }
			}`)
	})

	responseStruct, response, err := client.Wallets.Get(context.Background(), "dummy")
	testGeneralError(t, "Wallets.Get", err)
	testResponseUrl(t, "Wallets.Get", response, "/api/wallets/dummy")
	testResponseStruct(t, "Wallets.Get", responseStruct, &GetWallet{
		Data: Wallet{
			Address:    "dummy",
			PublicKey:  "dummy",
			Balance:    1000000000,
			IsDelegate: false,
		},
	})
}

// Get all transactions for the given wallet.
func TestWalletsService_Transactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummy",
			      "blockId": "dummy",
			      "type": 0,
			      "amount": 10000000,
			      "fee": 10000000,
			      "sender": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
			      "timestamp": {
			        "epoch": 40505460,
			        "unix": 1530606660,
			        "human": "2018-07-03T08:31:00Z"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.Transactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/dummy/transactions?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:            "dummy",
			BlockId:       "dummy",
			Type:          0,
			Amount:        10000000,
			Fee:           10000000,
			Sender:        "dummy",
			Recipient:     "dummy",
			Signature:     "dummy",
			VendorField:   "dummy",
			Confirmations: 10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
		}},
	})
}

// Get all transactions sent by the given wallet.
func TestWalletsService_SentTransactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions/sent", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions/sent?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummy",
			      "blockId": "dummy",
			      "type": 0,
			      "amount": 10000000,
			      "fee": 10000000,
			      "sender": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
			      "timestamp": {
			        "epoch": 40505460,
			        "unix": 1530606660,
			        "human": "2018-07-03T08:31:00Z"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.SentTransactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions/sent")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions/sent?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:            "dummy",
			BlockId:       "dummy",
			Type:          0,
			Amount:        10000000,
			Fee:           10000000,
			Sender:        "dummy",
			Recipient:     "dummy",
			Signature:     "dummy",
			VendorField:   "dummy",
			Confirmations: 10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
		}},
	})
}

// Get all transactions received by the given wallet.
func TestWalletsService_ReceivedTransaction(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions/received", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions/received?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions/received?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions/received?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummy",
			      "blockId": "dummy",
			      "type": 0,
			      "amount": 10000000,
			      "fee": 10000000,
			      "sender": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
			      "timestamp": {
			        "epoch": 40505460,
			        "unix": 1530606660,
			        "human": "2018-07-03T08:31:00Z"
			      }
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.ReceivedTransactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions/received")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/dummy/transactions/received?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions/received?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions/received?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:            "dummy",
			BlockId:       "dummy",
			Type:          0,
			Amount:        10000000,
			Fee:           10000000,
			Sender:        "dummy",
			Recipient:     "dummy",
			Signature:     "dummy",
			VendorField:   "dummy",
			Confirmations: 10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
		}},
	})
}

// Get all votes by the given wallet.
func TestWalletsService_Votes(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/votes", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/votes?page=1&limit=1",
			    "first": "/api/wallets/dummy/votes?page=1&limit=1",
			    "last": "/api/wallets/dummy/votes?page=1&limit=1"
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
	responseStruct, response, err := client.Wallets.Votes(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Votes", err)
	testResponseUrl(t, "Wallets.Votes", response, "/api/wallets/dummy/votes")
	testResponseStruct(t, "Wallets.Votes", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/dummy/votes?page=1&limit=1",
			First:      "/api/wallets/dummy/votes?page=1&limit=1",
			Last:       "/api/wallets/dummy/votes?page=1&limit=1",
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

// Filter all wallets by the given criteria.
func TestWalletsService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/search?page=1&limit=1",
			    "first": "/api/wallets/search?page=1&limit=1",
			    "last": "/api/wallets/search?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "balance": 1000000000,
			      "isDelegate": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &WalletsSearchRequest{Address: "dummy"}
	responseStruct, response, err := client.Wallets.Search(context.Background(), query, body)
	testGeneralError(t, "Wallets.Search", err)
	testResponseUrl(t, "Wallets.Search", response, "/api/wallets/search")
	testResponseStruct(t, "Wallets.Search", responseStruct, &Wallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/search?page=1&limit=1",
			First:      "/api/wallets/search?page=1&limit=1",
			Last:       "/api/wallets/search?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:    "dummy",
			PublicKey:  "dummy",
			Balance:    1000000000,
			IsDelegate: false,
		}},
	})
}

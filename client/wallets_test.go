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
			      "nonce": "1",
			      "balance": "1000000000",
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
			Nonce:      1,
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
			      "nonce": "1",
			      "balance": "1000000000",
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
			Nonce:      1,
			Balance:    1000000000,
			IsDelegate: false,
		}},
	})
}

// Get a wallet by the given id. (address, publicKey and username are valid)
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
		      "nonce": "1",
			    "balance": "1000000000",
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
			Nonce:      1,
			Balance:    1000000000,
			IsDelegate: false,
		},
	})
}

// Get all locks for the given wallet.
func TestWalletsService_Locks(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/locks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/locks?page=1&limit=1",
			    "first": "/api/wallets/dummy/locks?page=1&limit=1",
			    "last": "/api/wallets/dummy/locks?page=1&limit=1"
			  },
			  "data": [
			    {
			      "lockId": "dummy",
			      "amount": "1",
			      "secretHash": "dummySecretHash",
			      "senderPublicKey": "dummyPublicKey",
			      "recipientId": "dummyRecipient",
			      "timestamp": {
			        "epoch": 81911280,
			        "unix": 1572012480,
			        "human": "2019-10-25T14:08:00.000Z"
			      },
			      "expirationType": 2,
			      "expirationValue": 6000000,
			      "vendorField": "dummyVendorField"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Wallets.Locks(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Locks", err)
	testResponseUrl(t, "Wallets.Locks", response, "/api/wallets/dummy/locks")
	testResponseStruct(t, "Wallets.Locks", responseStruct, &Locks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/wallets/dummy/locks?page=1&limit=1",
			First:      "/api/wallets/dummy/locks?page=1&limit=1",
			Last:       "/api/wallets/dummy/locks?page=1&limit=1",
		},
		Data: []Lock{{
			LockId:          "dummy",
			Amount:          1,
			SecretHash:      "dummySecretHash",
			SenderPublicKey: "dummyPublicKey",
			RecipientId:     "dummyRecipient",
			Timestamp: Timestamp{
				Epoch: 81911280,
				Unix:  1572012480,
				Human: "2019-10-25T14:08:00.000Z",
			},
			ExpirationType:  2,
			ExpirationValue: 6000000,
			VendorField:     "dummyVendorField",
		}},
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
			      "typeGroup": 1,
			      "amount": "10000000",
			      "fee": "10000000",
			      "sender": "dummy",
			      "senderPublicKey": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
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
			Confirmations:   10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
			Nonce: 1,
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
			      "typeGroup": 1,
			      "amount": "10000000",
			      "fee": "10000000",
			      "sender": "dummy",
			      "senderPublicKey": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
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
			Confirmations:   10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
			Nonce: 1,
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
			      "typeGroup": 1,
			      "amount": "10000000",
			      "fee": "10000000",
			      "sender": "dummy",
			      "senderPublicKey": "dummy",
			      "recipient": "dummy",
			      "signature": "dummy",
			      "vendorField": "dummy",
			      "confirmations": 10,
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
			Confirmations:   10,
			Timestamp: Timestamp{
				Epoch: 40505460,
				Unix:  1530606660,
				Human: "2018-07-03T08:31:00Z",
			},
			Nonce: 1,
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
			      "typeGroup": 1,
			      "amount": "0",
			      "fee": "100000000",
			      "sender": "dummy",
			      "senderPublicKey": "dummy",
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
			      },
			      "nonce": "1"
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
			Id:              "dummy",
			BlockId:         "dummy",
			Type:            3,
			TypeGroup:       1,
			Amount:          0,
			Fee:             100000000,
			Sender:          "dummy",
			SenderPublicKey: "dummy",
			Recipient:       "dummy",
			Signature:       "dummy",
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
			Nonce: 1,
		}},
	})
}

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

// Get all transactions.
func TestTransactionsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/transactions?page=1&limit=1",
			    "first": "/api/transactions?page=1&limit=1",
			    "last": "/api/transactions?page=1&limit=1"
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
	responseStruct, response, err := client.Transactions.List(context.Background(), query)
	testGeneralError(t, "Transactions.List", err)
	testResponseUrl(t, "Transactions.List", response, "/api/transactions")
	testResponseStruct(t, "Transactions.List", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/transactions?page=1&limit=1",
			First:      "/api/transactions?page=1&limit=1",
			Last:       "/api/transactions?page=1&limit=1",
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

// Create a new transaction.
func TestTransactionsService_Create(t *testing.T) {
	t.Skip("test not ready")
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		testJsonPayload(t, request, values{"limit": 1})
		fmt.Fprint(writer,
			`{
			  "data": {
			    "accept": [
			    	"dummy"
			    ],
			    "excess": [],
			    "invalid": []
			  }
			}`)
	})

	body := &CreateTransactionRequest{
		Transactions: []Transaction{{
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
	}
	responseStruct, response, err := client.Transactions.Create(context.Background(), body)
	testGeneralError(t, "Transactions.Create", err)
	testResponseUrl(t, "Transactions.Create", response, "/api/transactions")
	testResponseStruct(t, "Transactions.Create", responseStruct, &GetCreateTransaction{
		Data: CreateTransaction{
			Accept: []string{
				"dummy",
			},
			Excess:  []string{},
			Invalid: []string{},
		},
	})
}

// Get a transaction by the given id.
func TestTransactionsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
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
			}`)
	})

	responseStruct, response, err := client.Transactions.Get(context.Background(), "dummy")
	testGeneralError(t, "Transactions.Get", err)
	testResponseUrl(t, "Transactions.Get", response, "/api/transactions/dummy")
	testResponseStruct(t, "Transactions.Get", responseStruct, &GetTransaction{
		Data: Transaction{
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
		},
	})
}

// Get all unconfirmed transactions.
func TestTransactionsService_ListUnconfirmed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/unconfirmed", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/transactions/unconfirmed?page=1&limit=1",
			    "first": "/api/transactions/unconfirmed?page=1&limit=1",
			    "last": "/api/transactions/unconfirmed?page=1&limit=1"
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
	responseStruct, response, err := client.Transactions.ListUnconfirmed(context.Background(), query)
	testGeneralError(t, "Transactions.ListUnconfirmed", err)
	testResponseUrl(t, "Transactions.ListUnconfirmed", response, "/api/transactions/unconfirmed")
	testResponseStruct(t, "Transactions.ListUnconfirmed", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/transactions/unconfirmed?page=1&limit=1",
			First:      "/api/transactions/unconfirmed?page=1&limit=1",
			Last:       "/api/transactions/unconfirmed?page=1&limit=1",
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

// Get an unconfirmed transaction by the given id.
func TestTransactionsService_GetUnconfirmed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/unconfirmed/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
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
			}`)
	})

	responseStruct, response, err := client.Transactions.GetUnconfirmed(context.Background(), "dummy")
	testGeneralError(t, "Transactions.GetUnconfirmed", err)
	testResponseUrl(t, "Transactions.GetUnconfirmed", response, "/api/transactions/unconfirmed/dummy")
	testResponseStruct(t, "Transactions.GetUnconfirmed", responseStruct, &GetTransaction{
		Data: Transaction{
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
		},
	})
}

// Filter all transactions by the given criteria.
func TestTransactionsService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/transactions/search?page=1&limit=1",
			    "first": "/api/transactions/search?page=1&limit=1",
			    "last": "/api/transactions/search?page=1&limit=1"
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
	body := &TransactionsSearchRequest{Id: "dummy"}
	responseStruct, response, err := client.Transactions.Search(context.Background(), query, body)
	testGeneralError(t, "Transactions.Search", err)
	testResponseUrl(t, "Transactions.Search", response, "/api/transactions/search")
	testResponseStruct(t, "Transactions.Search", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/transactions/search?page=1&limit=1",
			First:      "/api/transactions/search?page=1&limit=1",
			Last:       "/api/transactions/search?page=1&limit=1",
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

// Get a list of valid transaction types.
func TestTransactionsService_Types(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/types", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "1": {
			      "Transfer": 0,
			      "SecondSignature": 1,
			      "DelegateRegistration": 2,
			      "Vote": 3,
			      "MultiSignature": 4,
			      "Ipfs": 5,
			      "MultiPayment": 6,
			      "DelegateResignation": 7,
			      "HtlcLock": 8,
			      "HtlcClaim": 9,
			      "HtlcRefund": 10
			    },
			    "2": {
			      "BusinessRegistration": 0,
			      "BusinessResignation": 1,
			      "BusinessUpdate": 2,
			      "BridgechainRegistration": 3,
			      "BridgechainResignation": 4,
			      "BridgechainUpdate": 5
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Transactions.Types(context.Background())
	testGeneralError(t, "Transactions.Types", err)
	testResponseUrl(t, "Transactions.Types", response, "/api/transactions/types")
	testResponseStruct(t, "Transactions.Types", responseStruct, &TransactionTypes{
		Data: map[string]TypeGroupTypes{
			"1": {
				"Transfer":             0,
				"SecondSignature":      1,
				"DelegateRegistration": 2,
				"Vote":                 3,
				"MultiSignature":       4,
				"Ipfs":                 5,
				"MultiPayment":         6,
				"DelegateResignation":  7,
				"HtlcLock":             8,
				"HtlcClaim":            9,
				"HtlcRefund":           10,
			},
			"2": {
				"BusinessRegistration":    0,
				"BusinessResignation":     1,
				"BusinessUpdate":          2,
				"BridgechainRegistration": 3,
				"BridgechainResignation":  4,
				"BridgechainUpdate":       5,
			},
		},
	})
}

// Get a list of static transaction fees.
func TestTransactionsService_Fees(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/fees", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "transfer": 10000000,
			    "secondSignature": 500000000,
			    "delegateRegistration": 2500000000,
			    "vote": 100000000,
			    "multiSignature": 500000000,
			    "ipfs": 500000000,
			    "multiPayment": 10000000,
			    "delegateResignation": 2500000000,
			    "htlcLock": 10000000,
			    "htlcClaim": 0,
			    "htlcRefund": 0
			  }
			}`)
	})

	responseStruct, response, err := client.Transactions.Fees(context.Background())
	testGeneralError(t, "Transactions.Fees", err)
	testResponseUrl(t, "Transactions.Fees", response, "/api/transactions/fees")
	testResponseStruct(t, "Transactions.Fees", responseStruct, &TransactionFees{
		Data: map[string]uint32{
			"transfer":             10000000,
			"secondSignature":      500000000,
			"delegateRegistration": 2500000000,
			"vote":                 100000000,
			"multiSignature":       500000000,
			"ipfs":                 500000000,
			"multiPayment":         10000000,
			"delegateResignation":  2500000000,
			"htlcLock":             10000000,
			"htlcClaim":            0,
			"htlcRefund":           0,
		},
	})
}

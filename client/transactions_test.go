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
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "10000000",
			      "gas": "21000",
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
			        "cumulativeGasUsed": 21000,
			        "gasRefunded": 0,
			        "gasUsed": 21000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &TransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Transactions.List(context.Background(), query)
	testGeneralError(t, "Transactions.List", err)
	testResponseUrl(t, "Transactions.List", response, "/api/transactions")
	testResponseStruct(t, "Transactions.List", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/transactions?page=1&limit=1",
			First:      "/api/transactions?page=1&limit=1",
			Last:       "/api/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
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
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		}},
	})
}

// Get the transactions host's pool configuration.
func TestTransactionsService_Configuration(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/configuration", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "blockNumber": 23197739,
			    "core": {
			      "version": "0.0.1-evm.53"
			    },
			    "transactionPool": {
			      "maxTransactionAge": 2700,
			      "maxTransactionBytes": 128000,
			      "maxTransactionsInPool": 15000,
			      "maxTransactionsPerRequest": 40,
			      "maxTransactionsPerSender": 150
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Transactions.Configuration(context.Background())
	testGeneralError(t, "Transactions.Configuration", err)
	testResponseUrl(t, "Transactions.Configuration", response, "/tx/api/configuration")
	testResponseStruct(t, "Transactions.Configuration", responseStruct, &GetTransactionConfiguration{
		Data: TransactionConfiguration{
			Core: NodeCore{
				Version: "0.0.1-evm.53",
			},
			BlockNumber: 23197739,
			TransactionPool: TransactionConfigurationPool{
				MaxTransactionAge:         2700,
				MaxTransactionBytes:       128000,
				MaxTransactionsInPool:     15000,
				MaxTransactionsPerRequest: 40,
				MaxTransactionsPerSender:  150,
			},
		},
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
		Transactions: []string{"dummy"},
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
			    "hash": "dummy",
			    "blockHash": "dummy",
			    "value": "10000000",
			    "gas": "21000",
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
			      "cumulativeGasUsed": 21000,
			      "gasRefunded": 0,
			      "gasUsed": 21000,
			      "status": 1
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Transactions.Get(context.Background(), "dummy", nil)
	testGeneralError(t, "Transactions.Get", err)
	testResponseUrl(t, "Transactions.Get", response, "/api/transactions/dummy")
	testResponseStruct(t, "Transactions.Get", responseStruct, &GetTransaction{
		Data: Transaction{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
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
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
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
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "10000000",
			      "gas": "21000",
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
			        "cumulativeGasUsed": 21000,
			        "gasRefunded": 0,
			        "gasUsed": 21000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &UnconfirmedTransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Transactions.ListUnconfirmed(context.Background(), query)
	testGeneralError(t, "Transactions.ListUnconfirmed", err)
	testResponseUrl(t, "Transactions.ListUnconfirmed", response, "/api/transactions/unconfirmed")
	testResponseStruct(t, "Transactions.ListUnconfirmed", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/transactions/unconfirmed?page=1&limit=1",
			First:      "/api/transactions/unconfirmed?page=1&limit=1",
			Last:       "/api/transactions/unconfirmed?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
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
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
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
			    "hash": "dummy",
			    "blockHash": "dummy",
			    "value": "10000000",
			    "gas": "21000",
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
			      "cumulativeGasUsed": 21000,
			      "gasRefunded": 0,
			      "gasUsed": 21000,
			      "status": 1
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Transactions.GetUnconfirmed(context.Background(), "dummy")
	testGeneralError(t, "Transactions.GetUnconfirmed", err)
	testResponseUrl(t, "Transactions.GetUnconfirmed", response, "/api/transactions/unconfirmed/dummy")
	testResponseStruct(t, "Transactions.GetUnconfirmed", responseStruct, &GetTransaction{
		Data: Transaction{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
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
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		},
	})
}

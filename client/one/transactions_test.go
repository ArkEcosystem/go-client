// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all accounts.
func TestTransactionsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "transactions": [
			    {
			      "id": "dummy",
			      "blockid": "dummy",
			      "type": 0,
			      "timestamp": 10,
			      "amount": 10,
			      "fee": 10,
			      "vendorField": "dummy",
			      "senderId": "dummy",
			      "recipientId": "dummy",
			      "senderPublicKey": "dummy",
			      "signature": "dummy",
			      "signSignature": "dummy",
			      "asset": {},
			      "confirmations": 10
			    }
			  ]
			}`)
	})

	query := &GetTransactionsQuery{Limit: 1}
	responseStruct, response, err := client.Transactions.List(context.Background(), query)
	testGeneralError(t, "Transactions.List", err)
	testResponseUrl(t, "Transactions.List", response, "/api/transactions")
	testResponseStruct(t, "Transactions.List", responseStruct, &Transactions{
		Success: true,
		Transactions: []Transaction{{
			Id:              "dummy",
			BlockId:         "dummy",
			Type:            0,
			Timestamp:       10,
			Amount:          10,
			Fee:             10,
			VendorField:     "dummy",
			SenderId:        "dummy",
			RecipientId:     "dummy",
			SenderPublicKey: "dummy",
			Signature:       "dummy",
			SignSignature:   "dummy",
			Asset:           map[string]string{},
			Confirmations:   10,
		}},
	})
}

// Get a transaction by the given id.
func TestTransactionsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/get", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
		      "id": "dummy",
		      "blockid": "dummy",
		      "type": 0,
		      "timestamp": 10,
		      "amount": 10,
		      "fee": 10,
		      "vendorField": "dummy",
		      "senderId": "dummy",
		      "recipientId": "dummy",
		      "senderPublicKey": "dummy",
		      "signature": "dummy",
		      "signSignature": "dummy",
		      "asset": {},
		      "confirmations": 10
			}`)
	})

	responseStruct, response, err := client.Transactions.Get(context.Background(), "dummy")
	testGeneralError(t, "Transactions.Get", err)
	testResponseUrl(t, "Transactions.Get", response, "/api/transactions/get")
	testResponseStruct(t, "Transactions.Get", responseStruct, &Transaction{
		Success:         true,
		Id:              "dummy",
		BlockId:         "dummy",
		Type:            0,
		Timestamp:       10,
		Amount:          10,
		Fee:             10,
		VendorField:     "dummy",
		SenderId:        "dummy",
		RecipientId:     "dummy",
		SenderPublicKey: "dummy",
		Signature:       "dummy",
		SignSignature:   "dummy",
		Asset:           map[string]string{},
		Confirmations:   10,
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
			  "success": true,
			  "transactions": [
			    {
			      "id": "dummy",
			      "blockid": "dummy",
			      "type": 0,
			      "timestamp": 10,
			      "amount": 10,
			      "fee": 10,
			      "vendorField": "dummy",
			      "senderId": "dummy",
			      "recipientId": "dummy",
			      "senderPublicKey": "dummy",
			      "signature": "dummy",
			      "signSignature": "dummy",
			      "asset": {},
			      "confirmations": 10
			    }
			  ]
			}`)
	})

	query := &GetUnconfirmedTransactionsQuery{SenderPublicKey: "dummy"}
	responseStruct, response, err := client.Transactions.ListUnconfirmed(context.Background(), query)
	testGeneralError(t, "Transactions.ListUnconfirmed", err)
	testResponseUrl(t, "Transactions.ListUnconfirmed", response, "/api/transactions/unconfirmed")
	testResponseStruct(t, "Transactions.ListUnconfirmed", responseStruct, &Transactions{
		Success: true,
		Transactions: []Transaction{{
			Id:              "dummy",
			BlockId:         "dummy",
			Type:            0,
			Timestamp:       10,
			Amount:          10,
			Fee:             10,
			VendorField:     "dummy",
			SenderId:        "dummy",
			RecipientId:     "dummy",
			SenderPublicKey: "dummy",
			Signature:       "dummy",
			SignSignature:   "dummy",
			Asset:           map[string]string{},
			Confirmations:   10,
		}},
	})
}

// Get an unconfirmed transaction by the given id.
func TestTransactionsService_GetUnconfirmed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/transactions/unconfirmed/get", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
		      "id": "dummy",
		      "blockid": "dummy",
		      "type": 0,
		      "timestamp": 10,
		      "amount": 10,
		      "fee": 10,
		      "vendorField": "dummy",
		      "senderId": "dummy",
		      "recipientId": "dummy",
		      "senderPublicKey": "dummy",
		      "signature": "dummy",
		      "signSignature": "dummy",
		      "asset": {},
		      "confirmations": 10
			}`)
	})

	responseStruct, response, err := client.Transactions.GetUnconfirmed(context.Background(), "dummy")
	testGeneralError(t, "Transactions.GetUnconfirmed", err)
	testResponseUrl(t, "Transactions.GetUnconfirmed", response, "/api/transactions/unconfirmed/get")
	testResponseStruct(t, "Transactions.GetUnconfirmed", responseStruct, &Transaction{
		Success:         true,
		Id:              "dummy",
		BlockId:         "dummy",
		Type:            0,
		Timestamp:       10,
		Amount:          10,
		Fee:             10,
		VendorField:     "dummy",
		SenderId:        "dummy",
		RecipientId:     "dummy",
		SenderPublicKey: "dummy",
		Signature:       "dummy",
		SignSignature:   "dummy",
		Asset:           map[string]string{},
		Confirmations:   10,
	})
}

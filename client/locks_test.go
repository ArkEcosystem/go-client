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

// Get all locks.
func TestLocksService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/locks", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/locks?page=1&limit=1",
			    "first": "/api/locks?page=1&limit=1",
			    "last": "/api/locks?page=1&limit=1"
			  },
			  "data": [
			    {
			      "lockId": "dummyLockId",
			      "amount": "1",
			      "secretHash": "dummySecretHash",
			      "senderPublicKey": "dummySenderPublicKey",
			      "recipientId": "dummyRecipientId",
			      "timestamp": {
			        "epoch": 81911280,
			        "unix": 1572012480,
			        "human": "2019-10-25T14:08:00.000Z"
			      },
			      "expirationType": 2,
			      "expirationValue": 6000000,
			      "vendorField": "dummyVendorField",
			      "isExpired": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Locks.List(context.Background(), query)
	testGeneralError(t, "Locks.List", err)
	testResponseUrl(t, "Locks.List", response, "/api/locks")
	testResponseStruct(t, "Locks.List", responseStruct, &Locks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/locks?page=1&limit=1",
			First:      "/api/locks?page=1&limit=1",
			Last:       "/api/locks?page=1&limit=1",
		},
		Data: []Lock{{
			LockId:          "dummyLockId",
			Amount:          1,
			SecretHash:      "dummySecretHash",
			SenderPublicKey: "dummySenderPublicKey",
			RecipientId:     "dummyRecipientId",
			Timestamp: Timestamp{
				Epoch: 81911280,
				Unix:  1572012480,
				Human: "2019-10-25T14:08:00.000Z",
			},
			ExpirationType:  2,
			ExpirationValue: 6000000,
			VendorField:     "dummyVendorField",
			IsExpired:       false,
		}},
	})
}

// Get a lock by the given id
func TestLocksService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/locks/dummyLockId", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "lockId": "dummyLockId",
			    "amount": "1",
			    "secretHash": "dummySecretHash",
			    "senderPublicKey": "dummySenderPublicKey",
			    "recipientId": "dummyRecipientId",
			    "timestamp": {
			      "epoch": 81911280,
			      "unix": 1572012480,
			      "human": "2019-10-25T14:08:00.000Z"
			    },
			    "expirationType": 2,
			    "expirationValue": 6000000,
			    "vendorField": "dummyVendorField",
			    "isExpired": false
			  }
			}`)
	})

	responseStruct, response, err := client.Locks.Get(context.Background(), "dummyLockId")
	testGeneralError(t, "Locks.Get", err)
	testResponseUrl(t, "Locks.Get", response, "/locks/dummyLockId")
	testResponseStruct(t, "Locks.Get", responseStruct, &GetLock{
		Data: Lock{
			LockId:          "dummyLockId",
			Amount:          1,
			SecretHash:      "dummySecretHash",
			SenderPublicKey: "dummySenderPublicKey",
			RecipientId:     "dummyRecipientId",
			Timestamp: Timestamp{
				Epoch: 81911280,
				Unix:  1572012480,
				Human: "2019-10-25T14:08:00.000Z",
			},
			ExpirationType:  2,
			ExpirationValue: 6000000,
			VendorField:     "dummyVendorField",
			IsExpired:       false,
		},
	})
}

// Filter all locks by the given criteria.
func TestLocksService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/locks/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/locks/search?page=1&limit=1",
			    "first": "/api/locks/search?page=1&limit=1",
			    "last": "/api/locks/search?page=1&limit=1"
			  },
			  "data": [
			    {
			      "lockId": "dummyLockId",
			      "amount": "1",
			      "secretHash": "dummySecretHash",
			      "senderPublicKey": "dummySenderPublicKey",
			      "recipientId": "dummyRecipientId",
			      "timestamp": {
			        "epoch": 81911280,
			        "unix": 1572012480,
			        "human": "2019-10-25T14:08:00.000Z"
			      },
			      "expirationType": 2,
			      "expirationValue": 6000000,
			      "vendorField": "dummyVendorField",
			      "isExpired": false
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &LocksSearchRequest{LockId: "dummyLockId"}
	responseStruct, response, err := client.Locks.Search(context.Background(), query, body)
	testGeneralError(t, "Locks.Search", err)
	testResponseUrl(t, "Locks.Search", response, "/api/locks/search")
	testResponseStruct(t, "Locks.Search", responseStruct, &Locks{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/locks/search?page=1&limit=1",
			First:      "/api/locks/search?page=1&limit=1",
			Last:       "/api/locks/search?page=1&limit=1",
		},
		Data: []Lock{{
			LockId:          "dummyLockId",
			Amount:          1,
			SecretHash:      "dummySecretHash",
			SenderPublicKey: "dummySenderPublicKey",
			RecipientId:     "dummyRecipientId",
			Timestamp: Timestamp{
				Epoch: 81911280,
				Unix:  1572012480,
				Human: "2019-10-25T14:08:00.000Z",
			},
			ExpirationType:  2,
			ExpirationValue: 6000000,
			VendorField:     "dummyVendorField",
			IsExpired:       false,
		}},
	})
}

// Retrieve transactions by the given lock ids.
func TestLocksService_Unlocked(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/locks/unlocked", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/locks/unlocked?page=1&limit=1",
			    "first": "/api/locks/unlocked?page=1&limit=1",
			    "last": "/api/locks/unlocked?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummyId",
			      "blockId": "dummyBlockId",
			      "version": 2,
			      "type": 0,
			      "typeGroup": 1,
			      "amount": "0",
			      "fee": "0",
			      "sender": "dummySender",
			      "senderPublicKey": "dummySenderPublicKey",
			      "recipient": "dummyRecipient",
			      "signature": "dummySignature",
			      "asset": {
			        "claim": {
			          "lockTransactionId": "dummyLockTransactionId",
			          "unlockSecret": "dummyUnlockSecret"
			        }
			      },
			      "vendorField": "dummy",
			      "confirmations": 3,
			      "timestamp": {
			        "epoch": 82354848,
			        "unix": 1572456048,
			        "human": "2019-10-30T17:20:48.000Z"
			      },
			      "nonce": "1"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &LocksUnlockedRequest{
		Ids: []string{
			"dummyLockId",
		},
	}
	responseStruct, response, err := client.Locks.Unlocked(context.Background(), query, body)
	testGeneralError(t, "Locks.Unlocked", err)
	testResponseUrl(t, "Locks.Unlocked", response, "/api/locks/unlocked")
	testResponseStruct(t, "Locks.Unlocked", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/locks/unlocked?page=1&limit=1",
			First:      "/api/locks/unlocked?page=1&limit=1",
			Last:       "/api/locks/unlocked?page=1&limit=1",
		},
		Data: []Transaction{{
			Id:              "dummyId",
			BlockId:         "dummyBlockId",
			Version:         2,
			Type:            0,
			TypeGroup:       1,
			Amount:          0,
			Fee:             0,
			Sender:          "dummySender",
			SenderPublicKey: "dummySenderPublicKey",
			Recipient:       "dummyRecipient",
			Signature:       "dummySignature",
			Asset: &TransactionAsset{
				Claim: &ClaimAsset{
					LockTransactionId: "dummyLockTransactionId",
					UnlockSecret:      "dummyUnlockSecret",
				},
			},
			Confirmations: 3,
			VendorField:   "dummy",
			Timestamp: Timestamp{
				Epoch: 82354848,
				Unix:  1572456048,
				Human: "2019-10-30T17:20:48.000Z",
			},
			Nonce: 1,
		}},
	})
}

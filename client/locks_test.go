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
			    "last": "/api/locks?page=7&limit=1"
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
			      "vendorField": "dummyVendorField"
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
			Amount:          "1",
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
			    "vendorField": "dummyVendorField"
			  }
			}`)
	})

	responseStruct, response, err := client.Locks.Get(context.Background(), "dummyLockId")
	testGeneralError(t, "Locks.Get", err)
	testResponseUrl(t, "Locks.Get", response, "/locks/dummyLockId")
	testResponseStruct(t, "Locks.Get", responseStruct, &GetLock{
		Data: Lock{
			LockId:          "dummyLockId",
			Amount:          "1",
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
			    "self": "/api/locks?page=1&limit=1",
			    "first": "/api/locks?page=1&limit=1",
			    "last": "/api/locks?page=7&limit=1"
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
			      "vendorField": "dummyVendorField"
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
			Self:       "/api/locks?page=1&limit=1",
			First:      "/api/locks?page=1&limit=1",
			Last:       "/api/locks?page=1&limit=1",
		},
		Data: []Lock{{
			LockId:          "dummyLockId",
			Amount:          "1",
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
		}},
	})
}

// Filter all locks by the given ids.
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
			    "self": "/api/locks?page=1&limit=1",
			    "first": "/api/locks?page=1&limit=1",
			    "last": "/api/locks?page=7&limit=1"
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
			      "vendorField": "dummyVendorField"
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
	testResponseStruct(t, "Locks.Unlocked", responseStruct, &Locks{
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
			Amount:          "1",
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
		}},
	})
}

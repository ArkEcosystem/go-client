// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all accounts.
func TestAccountsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/getAllAccounts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
		  "accounts": [{
		    "address": "dummy",
		    "publicKey": "dummy",
		    "secondPublicKey": "dummy",
		    "username": "dummy",
		    "balance": "dummy",
		    "unconfirmedBalance": "dummy",
		    "multisignatures": [],
		    "u_multisignatures": [],
		    "unconfirmedSignature": 0,
		    "secondSignature": 0
		  }],
		  "success": true
		}`)
	})

	responseStruct, response, err := client.Accounts.List(context.Background())
	testGeneralError(t, "Accounts.List", err)
	testResponseUrl(t, "Accounts.List", response, "/api/accounts/getAllAccounts")
	testResponseStruct(t, "Accounts.List", responseStruct, &Accounts{
		Success: true,
		Accounts: []Account{{
			Address:                    "dummy",
			PublicKey:                  "dummy",
			SecondPublicKey:            "dummy",
			Username:                   "dummy",
			Balance:                    "dummy",
			UnconfirmedBalance:         "dummy",
			Multisignatures:            []string{},
			UnconfirmedMultisignatures: []string{},
			UnconfirmedSignature:       0,
			SecondSignature:            0,
		}},
	})
}

// Get a account by the given address.
func TestAccountsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "account": {
			    "address": "dummy",
			    "publicKey": "dummy",
			    "secondPublicKey": "dummy",
			    "username": "dummy",
			    "balance": "dummy",
			    "unconfirmedBalance": "dummy",
			    "multisignatures": [],
			    "u_multisignatures": [],
			    "unconfirmedSignature": 0,
			    "secondSignature": 0
			  },
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Accounts.Get(context.Background(), "dummy")
	testGeneralError(t, "Accounts.Get", err)
	testResponseUrl(t, "Accounts.Get", response, "/api/accounts")
	testResponseStruct(t, "Accounts.Get", responseStruct, &AccountSingle{
		Success: true,
		Account: Account{
			Address:                    "dummy",
			PublicKey:                  "dummy",
			SecondPublicKey:            "dummy",
			Username:                   "dummy",
			Balance:                    "dummy",
			UnconfirmedBalance:         "dummy",
			Multisignatures:            []string{},
			UnconfirmedMultisignatures: []string{},
			UnconfirmedSignature:       0,
			SecondSignature:            0,
		},
	})
}

// Count all accounts.
func TestAccountsService_Count(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/count", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "count": 3,
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Accounts.Count(context.Background())
	testGeneralError(t, "Accounts.Count", err)
	testResponseUrl(t, "Accounts.Count", response, "/api/accounts/count")
	testResponseStruct(t, "Accounts.Count", responseStruct, &AccountsCount{
		Success: true,
		Count: 3,
	})
}

// Get a delegate by the given address.
func TestAccountsService_Delegate(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/delegates", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "delegates": [{
			    "username": "dummy",
			    "address": "dummy",
			    "publicKey": "dummy",
			    "vote": "dummy",
			    "producedblocks": 1,
			    "missedblocks": 2,
			    "rate": 3,
			    "approval": 0.5,
			    "productivity": 0.10
			  }],
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Accounts.Delegate(context.Background(), "dummy")
	testGeneralError(t, "Accounts.Delegate", err)
	testResponseUrl(t, "Accounts.Delegate", response, "/api/accounts/delegates")
	testResponseStruct(t, "Accounts.Delegate", responseStruct, &AccountDelegates{
		Success: true,
		Delegates: []Delegate {{
			Username: "dummy",
			Address: "dummy",
			PublicKey: "dummy",
			Vote: "dummy",
			ProducedBlocks: 1,
			MissedBlocks: 2,
			Rate: 3,
			Approval: 0.5,
			Productivity: 0.10,
		}},
	})
}

// Get the delegate registration fee.
func TestAccountsService_DelegateFee(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/delegates/fee", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "fee": 10,
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Accounts.DelegateFee(context.Background())
	testGeneralError(t, "Accounts.DelegateFee", err)
	testResponseUrl(t, "Accounts.DelegateFee", response, "/api/accounts/delegates/fee")
	testResponseStruct(t, "Accounts.DelegateFee", responseStruct, &DelegateFee{
		Success: true,
		Fee: 10,
	})
}

// Get the balance for an account by the given address.
func TestAccountsService_Balance(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/getBalance", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "balance": "10",
			  "unconfirmedBalance": "10",
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Accounts.Balance(context.Background(), "dummy")
	testGeneralError(t, "Accounts.Balance", err)
	testResponseUrl(t, "Accounts.Balance", response, "/api/accounts/getBalance")
	testResponseStruct(t, "Accounts.Balance", responseStruct, &AccountBalance{
		Success: true,
		Balance: "10",
		UnconfirmedBalance: "10",
	})
}

// Get the public key for an account by the given address.
func TestAccountsService_PublicKey(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/getPublicKey", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{"success":true,"publicKey":"dummy"}`)
	})

	responseStruct, response, err := client.Accounts.PublicKey(context.Background(), "dummy")
	testGeneralError(t, "Accounts.PublicKey", err)
	testResponseUrl(t, "Accounts.PublicKey", response, "/api/accounts/getPublicKey?address=dummy")
	testResponseStruct(t, "Accounts.PublicKey", responseStruct, &PublicKey{Success: true, PublicKey: "dummy"})
}

// Get all wallets sorted by balance in descending order.
func TestAccountsService_Top(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/top", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "accounts": [{
			    "address": "dummy",
			    "balance": "dummy",
			    "publicKey": "dummy"
			  }],
			  "success": true
			}`)
	})

	query := &TopQuery{Limit: 1}
	responseStruct, response, err := client.Accounts.Top(context.Background(), query)
	testGeneralError(t, "Accounts.Top", err)
	testResponseUrl(t, "Accounts.Top", response, "/api/accounts/top")
	testResponseStruct(t, "Accounts.Top", responseStruct, &AccountsTop{
		Success: true,
		Accounts: []Account{{
			Address:   "dummy",
			Balance:   "dummy",
			PublicKey: "dummy",
		}},
	})
}

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
	//
}

// Count all accounts.
func TestAccountsService_Count(t *testing.T) {
	//
}

// Get a delegate by the given address.
func TestAccountsService_Delegate(t *testing.T) {
	//
}

// Get the delegate registration fee.
func TestAccountsService_DelegateFee(t *testing.T) {
	//
}

// Get the balance for an account by the given address.
func TestAccountsService_Balance(t *testing.T) {
	//
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
	//
}

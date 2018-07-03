// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

// Get all accounts.
func TestAccountsService_List(t *testing.T) {
	//
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

	mux.HandleFunc("/accounts/getAllAccounts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{"success":true}`)
	})

	responseStruct, response, err := client.Accounts.List(context.Background())
	if err != nil {
		t.Errorf("Accounts.List returned error: %v", err)
	}

	assert := assert.New(t)
	assert.True(strings.Contains(response.Request.URL.String(), "/api/accounts/getAllAccounts"))
	assert.True(responseStruct.Success)
}

// Get all wallets sorted by balance in descending order.
func TestAccountsService_Top(t *testing.T) {
	//
}

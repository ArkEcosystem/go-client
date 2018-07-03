// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

type SuccessResponse struct {
	Success bool `json:"success,omitempty"`
}

// Get all accounts.
func TestOneAccountsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/getAllAccounts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{"success":true}`)
	})

	// responseModel := &SuccessResponse{}
	_, response, err := client.One_Accounts.List(context.Background(), &SuccessResponse{})
	if err != nil {
		t.Errorf("One_Accounts.List returned error: %v", err)
	}

	assert := assert.New(t)
	assert.True(strings.Contains(response.Request.URL.String(), "/api/accounts/getAllAccounts"))
	assert.True(responseModel.Success)
}

// Get a account by the given address.
func TestOneAccountsService_Get(t *testing.T) {
	//
}

// Count all accounts.
func TestOneAccountsService_Count(t *testing.T) {
	//
}

// Get a delegate by the given address.
func TestOneAccountsService_Delegate(t *testing.T) {
	//
}

// Get the delegate registration fee.
func TestOneAccountsService_DelegateFee(t *testing.T) {
	//
}

// Get the balance for an account by the given address.
func TestOneAccountsService_Balance(t *testing.T) {
	//
}

// Get the public key for an account by the given address.
func TestOneAccountsService_PublicKey(t *testing.T) {
	//
}

// Get all wallets sorted by balance in descending order.
func TestOneAccountsService_Top(t *testing.T) {
	//
}

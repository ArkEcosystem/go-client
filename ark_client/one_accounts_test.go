// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	// "encoding/json"
	"fmt"
	"net/http"
	// "reflect"
	"testing"
	"github.com/davecgh/go-spew/spew"
)

// Get all accounts.
func TestOneAccountsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/accounts/getAllAccounts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		spew.Dump(request)
		testFormValues(t, request, values{})
		fmt.Fprint(writer, `[{"id":1}]`)
	})

	response, _, err := client.One_Accounts.List(context.Background(), nil)
	// opt := &RepositoryListAllOptions{1}
	// response, _, err := client.One_Accounts.List(context.Background(), opt)
	if err != nil {
		t.Errorf("One_Accounts.List returned error: %v", err)
	}

	spew.Dump(response)
	// want := []*Repository{{ID: Int64(1)}}
	// if !reflect.DeepEqual(response, want) {
	// 	t.Errorf("One_Accounts.List returned %+v, want %+v", response, want)
	// }
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

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"net/http"
)

// TransactionsService handles communication with the transactions related
// methods of the Ark Core API - Version 2.
type TransactionsService Service

// Get all transactions.
func (s *TransactionsService) List(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions", query, nil)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Create a new transaction.
// TODO: add struct for request body
func (s *TransactionsService) Create(ctx context.Context) (*CreateTransaction, *http.Response, error) {
	var responseStruct *CreateTransaction
	resp, err := s.client.SendRequest(ctx, "POST", "transactions", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id int) (*Transaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/%v", id)

	var responseStruct *Transaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/unconfirmed", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id int) (*Transaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

	var responseStruct *Transaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all transactions by the given criteria.
func (s *TransactionsService) Search(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "POST", "transactions/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of valid transaction types.
func (s *TransactionsService) Types(ctx context.Context) (*TransactionTypes, *http.Response, error) {
	var responseStruct *TransactionTypes
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/types", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

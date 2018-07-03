// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"net/http"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type TransactionsService Service

// Get all accounts.
func (s *TransactionsService) List(ctx context.Context, query *GetTransactionsQuery) (interface{}, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &TransactionIdQuery{Id: id}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context, query *GetUnconfirmedTransactionsQuery) (interface{}, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &TransactionIdQuery{Id: id}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

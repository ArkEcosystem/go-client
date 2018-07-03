// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./requests/one"
	"./responses/one"
	"context"
	"net/http"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type One_TransactionsService Service

// Get all accounts.
func (s *One_TransactionsService) List(ctx context.Context, query *requests_one.GetTransactionsQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *One_TransactionsService) Get(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &requests_one.TransactionIdQuery{Id: id}

	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *One_TransactionsService) ListUnconfirmed(ctx context.Context, query *requests_one.GetUnconfirmedTransactionsQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *One_TransactionsService) GetUnconfirmed(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &requests_one.TransactionIdQuery{Id: id}

	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

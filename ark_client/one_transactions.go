// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"./structs/request"
	"context"
	"net/http"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type One_TransactionsService Service

// Get all accounts.
func (s *One_TransactionsService) List(ctx context.Context, query *request.GetTransactionsQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a transaction by the given id.
func (s *One_TransactionsService) Get(ctx context.Context, id string, model interface{}) (interface{}, *http.Response, error) {
	query := &request.TransactionIdQuery{Id: id}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/get", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all unconfirmed transactions.
func (s *One_TransactionsService) ListUnconfirmed(ctx context.Context, query *request.GetUnconfirmedTransactionsQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *One_TransactionsService) GetUnconfirmed(ctx context.Context, id string, model interface{}) (interface{}, *http.Response, error) {
	query := &request.TransactionIdQuery{Id: id}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "transactions/unconfirmed/get", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

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
type Two_TransactionsService Service

// Get all transactions.
func (s *Two_TransactionsService) List(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "transactions", query, nil)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Create a new transaction.
// TODO: Finish method
// func (s *Two_TransactionsService) Create(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post("transactions")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get a transaction by the given id.
func (s *Two_TransactionsService) Get(ctx context.Context, id int) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("transactions/%v", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *Two_TransactionsService) ListUnconfirmed(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "transactions/unconfirmed", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *Two_TransactionsService) GetUnconfirmed(ctx context.Context, id int) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all transactions by the given criteria.
// TODO: Finish method
// func (s *Two_TransactionsService) Search(ctx context.Context, query *requests_two.Pagination) (*http.Response, error) {
// 	resp, err := s.client.SendRequest(ctx, 2, "POST", "transactions/search", query, &responseStruct)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return responseStruct, resp, err
// }

// Get a list of valid transaction types.
func (s *Two_TransactionsService) Types(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "transactions/types", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	"fmt"
	"net/http"
)

// TransactionsService handles communication with the transactions related
// methods of the Ark Core API - Version 2.
type Two_TransactionsService Service

// Get all transactions.
func (s *Two_TransactionsService) List(ctx context.Context) (*http.Response, error) {
	return s.client.SendRequest(ctx, 2, "GET", "transactions", nil)
}

// Create a new transaction.
// func (s *Two_TransactionsService) Create(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post("transactions")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get a transaction by the given id.
func (s *Two_TransactionsService) Get(ctx context.Context, id int) (*http.Response, error) {
	uri := fmt.Sprintf("transactions/%v", id)

	return s.client.SendRequest(ctx, 2, "GET", uri, nil)
}

// Get all unconfirmed transactions.
func (s *Two_TransactionsService) ListUnconfirmed(ctx context.Context) (*http.Response, error) {
	return s.client.SendRequest(ctx, 2, "GET", "transactions/unconfirmed", nil)
}

// Get an unconfirmed transaction by the given id.
func (s *Two_TransactionsService) GetUnconfirmed(ctx context.Context, id int) (*http.Response, error) {
	uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

	return s.client.SendRequest(ctx, 2, "GET", uri, nil)
}

// Filter all transactions by the given criteria.
// func (s *Two_TransactionsService) Search(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post("transactions/search")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get a list of valid transaction types.
func (s *Two_TransactionsService) Types(ctx context.Context) (*http.Response, error) {
	return s.client.SendRequest(ctx, 2, "GET", "transactions/types", nil)
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "net/http"

    . "../types"
)

// TransactionsService handles communication with the transactions related
// methods of the Ark Core API - Version 2.
type TransactionsService Service

// Get all transactions.
func (s *TransactionsService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "transactions")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Create a new transaction.
// func (s *TransactionsService) Create(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post(s.Client.BaseURL.String() + "transactions")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "transactions/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "transactions/unconfirmed")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "transactions/unconfirmed/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Filter all transactions by the given criteria.
// func (s *TransactionsService) Search(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post(s.Client.BaseURL.String() + "transactions/search")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get a list of valid transaction types.
func (s *TransactionsService) Types(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "transactions/types")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

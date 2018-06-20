// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type One_TransactionsService Service

// Get all accounts.
func (s *One_TransactionsService) List(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(1, "GET", "transactions", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a transaction by the given id.
func (s *One_TransactionsService) Get(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(1, "GET", "transactions/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all unconfirmed transactions.
func (s *One_TransactionsService) ListUnconfirmed(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(1, "GET", "transactions/unconfirmed", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get an unconfirmed transaction by the given id.
func (s *One_TransactionsService) GetUnconfirmed(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(1, "GET", "transactions/unconfirmed/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}
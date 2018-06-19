// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "net/http"

    . "../types"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type TransactionsService Service

// Get all accounts.
func (s *TransactionsService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "api/transactions")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "api/transactions/get")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "api/transactions/unconfirmed")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "api/transactions/unconfirmed/get")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

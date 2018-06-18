// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type TransactionsService service

// Get all accounts.
func (s *TransactionsService) List(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/transactions", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/transactions/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/transactions/unconfirmed", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/transactions/unconfirmed/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

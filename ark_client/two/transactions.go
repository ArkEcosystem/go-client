// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "strings"
)

// TransactionsService handles communication with the transactions related
// methods of the Ark Core API - Version 2.
type TransactionsService service

// Get all transactions.
func (s *TransactionsService) List(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "transactions", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Create a new transaction.
func (s *TransactionsService) Create(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("POST", "transactions", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context) (*Response, error)
{
    uri := fmt.Sprintf("transactions/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "transactions/unconfirmed", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context) (*Response, error)
{
    uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Filter all transactions by the given criteria.
func (s *TransactionsService) Search(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("POST", "transactions/search", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a list of valid transaction types.
func (s *TransactionsService) Types(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "transactions/types", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

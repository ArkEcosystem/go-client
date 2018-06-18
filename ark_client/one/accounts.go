// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "fmt"
    "strings"
)

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type AccountsService service

// Get all accounts.
func (s *AccountsService) List(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/getAllAccounts", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a account by the given address.
func (s *AccountsService) Get(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Count all accounts.
func (s *AccountsService) Count(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/count", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a delegate by the given address.
func (s *AccountsService) Delegate(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/delegates", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the delegate registration fee.
func (s *AccountsService) DelegateFee(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/delegates/fee", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the balance for an account by the given address.
func (s *AccountsService) Balance(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/getBalance", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the public key for an account by the given address.
func (s *AccountsService) PublicKey(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/getPublicKey", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all wallets sorted by balance in descending order.
func (s *AccountsService) Top(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/accounts/top", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

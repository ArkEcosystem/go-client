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

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type AccountsService Service

// Get all accounts.
func (s *AccountsService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/getAllAccounts")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a account by the given address.
func (s *AccountsService) Get(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Count all accounts.
func (s *AccountsService) Count(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/count")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a delegate by the given address.
func (s *AccountsService) Delegate(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/delegates")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the delegate registration fee.
func (s *AccountsService) DelegateFee(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/delegates/fee")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the balance for an account by the given address.
func (s *AccountsService) Balance(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/getBalance")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the public key for an account by the given address.
func (s *AccountsService) PublicKey(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/getPublicKey")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all wallets sorted by balance in descending order.
func (s *AccountsService) Top(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/accounts/top")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

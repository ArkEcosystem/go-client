// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type One_AccountsService Service

// Get all accounts.
func (s *One_AccountsService) List(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/getAllAccounts", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a account by the given address.
func (s *One_AccountsService) Get(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Count all accounts.
func (s *One_AccountsService) Count(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/count", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a delegate by the given address.
func (s *One_AccountsService) Delegate(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/delegates", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the delegate registration fee.
func (s *One_AccountsService) DelegateFee(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/delegates/fee", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the balance for an account by the given address.
func (s *One_AccountsService) Balance(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/getBalance", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the public key for an account by the given address.
func (s *One_AccountsService) PublicKey(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/getPublicKey", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get all wallets sorted by balance in descending order.
func (s *One_AccountsService) Top(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "accounts/top", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

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

// WalletsService handles communication with the wallets related
// methods of the Ark Core API - Version 2.
type WalletsService Service

// Get all wallets.
func (s *WalletsService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("wallets")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all wallets sorted by balance in descending order.
func (s *WalletsService) Top(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("wallets/top")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a wallet by the given id.
func (s *WalletsService) Get(ctx context.Context) (*http.Response, error) {
    uri := fmt.Sprintf("wallets/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions for the given wallet.
func (s *WalletsService) Transactions(ctx context.Context) (*http.Response, error) {
    uri := fmt.Sprintf("wallets/%v/transactions", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions sent by the given wallet.
func (s *WalletsService) SentTransactions(ctx context.Context) (*http.Response, error) {
    uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions received by the given wallet.
func (s *WalletsService) ReceivedTransaction(ctx context.Context) (*http.Response, error) {
    uri := fmt.Sprintf("wallets/%v/transactions/received", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all votes by the given wallet.
func (s *WalletsService) Votes(ctx context.Context) (*http.Response, error) {
    uri := fmt.Sprintf("wallets/%v/votes", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Filter all wallets by the given criteria.
func (s *WalletsService) Search(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest("POST", "wallets/search", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

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

// WalletsService handles communication with the wallets related
// methods of the Ark Core API - Version 2.
type Two_WalletsService Service

// Get all wallets.
func (s *Two_WalletsService) List(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(2, "GET", "wallets", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all wallets sorted by balance in descending order.
func (s *Two_WalletsService) Top(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(2, "GET", "wallets/top", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a wallet by the given id.
func (s *Two_WalletsService) Get(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("wallets/%v", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all transactions for the given wallet.
func (s *Two_WalletsService) Transactions(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("wallets/%v/transactions", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all transactions sent by the given wallet.
func (s *Two_WalletsService) SentTransactions(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all transactions received by the given wallet.
func (s *Two_WalletsService) ReceivedTransaction(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("wallets/%v/transactions/received", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all votes by the given wallet.
func (s *Two_WalletsService) Votes(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("wallets/%v/votes", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Filter all wallets by the given criteria.
// func (s *Two_WalletsService) Search(ctx context.Context) (*http.Request, error) {
//     resp, err := s.Client.Client.Post("wallets/search")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

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
func (s *Two_WalletsService) List(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *Two_WalletsService) Top(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets/top", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a wallet by the given id.
func (s *Two_WalletsService) Get(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all transactions for the given wallet.
func (s *Two_WalletsService) Transactions(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all transactions sent by the given wallet.
func (s *Two_WalletsService) SentTransactions(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all transactions received by the given wallet.
func (s *Two_WalletsService) ReceivedTransaction(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/received", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all votes by the given wallet.
func (s *Two_WalletsService) Votes(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/votes", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Filter all wallets by the given criteria.
// func (s *Two_WalletsService) Search(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post("wallets/search")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

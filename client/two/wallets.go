// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"net/http"
)

// WalletsService handles communication with the wallets related
// methods of the Ark Core API - Version 2.
type WalletsService Service

// Get all wallets.
func (s *WalletsService) List(ctx context.Context, query *Pagination) (*Wallets, *http.Response, error) {
	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *WalletsService) Top(ctx context.Context, query *Pagination) (*Wallets, *http.Response, error) {
	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets/top", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a wallet by the given id.
func (s *WalletsService) Get(ctx context.Context, id int) (*Wallet, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v", id)

	var responseStruct *Wallet
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions for the given wallet.
func (s *WalletsService) Transactions(ctx context.Context, id int, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions sent by the given wallet.
func (s *WalletsService) SentTransactions(ctx context.Context, id int, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions received by the given wallet.
func (s *WalletsService) ReceivedTransaction(ctx context.Context, id int, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/received", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all votes by the given wallet.
func (s *WalletsService) Votes(ctx context.Context, id int, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/votes", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all wallets by the given criteria.
// func (s *WalletsService) Search(ctx context.Context, query *Pagination) (*http.Response, error) {
// 	resp, err := s.client.SendRequest(ctx, 2, "POST", "wallets/search", query, &responseStruct)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return responseStruct, resp, err
// }

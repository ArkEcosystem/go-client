// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"./structs/request"
	"context"
	"fmt"
	"net/http"
)

// WalletsService handles communication with the wallets related
// methods of the Ark Core API - Version 2.
type Two_WalletsService Service

// Get all wallets.
func (s *Two_WalletsService) List(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *Two_WalletsService) Top(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "wallets/top", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a wallet by the given id.
func (s *Two_WalletsService) Get(ctx context.Context, id int, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all transactions for the given wallet.
func (s *Two_WalletsService) Transactions(ctx context.Context, id int, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all transactions sent by the given wallet.
func (s *Two_WalletsService) SentTransactions(ctx context.Context, id int, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all transactions received by the given wallet.
func (s *Two_WalletsService) ReceivedTransaction(ctx context.Context, id int, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/received", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all votes by the given wallet.
func (s *Two_WalletsService) Votes(ctx context.Context, id int, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/votes", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Filter all wallets by the given criteria.
// func (s *Two_WalletsService) Search(ctx context.Context, query *request.Pagination) (*http.Response, error) {
// 	resp, err := s.client.SendRequest(ctx, 2, "POST", "wallets/search", query, &model)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return model, resp, err
// }

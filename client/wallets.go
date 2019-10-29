// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

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
	resp, err := s.client.SendRequest(ctx, "GET", "wallets", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *WalletsService) Top(ctx context.Context, query *Pagination) (*Wallets, *http.Response, error) {
	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, "GET", "wallets/top", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a wallet by the given id. (address, publicKey and username are valid)
func (s *WalletsService) Get(ctx context.Context, id string) (*GetWallet, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v", id)

	var responseStruct *GetWallet
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all locks for the given wallet.
func (s *WalletsService) Locks(ctx context.Context, id string, query *Pagination) (*Locks, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/locks", id)

	var responseStruct *Locks
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions for the given wallet.
func (s *WalletsService) Transactions(ctx context.Context, id string, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions sent by the given wallet.
func (s *WalletsService) SentTransactions(ctx context.Context, id string, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions received by the given wallet.
func (s *WalletsService) ReceivedTransactions(ctx context.Context, id string, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/transactions/received", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all votes by the given wallet.
func (s *WalletsService) Votes(ctx context.Context, id string, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("wallets/%v/votes", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all wallets by the given criteria.
func (s *WalletsService) Search(ctx context.Context, query *Pagination, body *WalletsSearchRequest) (*Wallets, *http.Response, error) {
	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, "POST", "wallets/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

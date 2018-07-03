// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"net/http"
)

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type AccountsService Service

// Get all accounts.
func (s *AccountsService) List(ctx context.Context) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getAllAccounts", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a account by the given address.
func (s *AccountsService) Get(ctx context.Context, address string) (*PublicKey, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Count all accounts.
func (s *AccountsService) Count(ctx context.Context) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/count", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a delegate by the given address.
func (s *AccountsService) Delegate(ctx context.Context, address string) (*PublicKey, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the delegate registration fee.
func (s *AccountsService) DelegateFee(ctx context.Context) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the balance for an account by the given address.
func (s *AccountsService) Balance(ctx context.Context, address string) (*PublicKey, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getBalance", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the public key for an account by the given address.
func (s *AccountsService) PublicKey(ctx context.Context, address string) (*PublicKey, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getPublicKey", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *AccountsService) Top(ctx context.Context, query *TopQuery) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/top", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

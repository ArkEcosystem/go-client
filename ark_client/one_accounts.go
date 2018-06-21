// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	"net/http"
)

type PublicKeyQuery struct {
	Address string `url:"address"`
}

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type One_AccountsService Service

// Get all accounts.
func (s *One_AccountsService) List(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getAllAccounts", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a account by the given address.
func (s *One_AccountsService) Get(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Count all accounts.
func (s *One_AccountsService) Count(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/count", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a delegate by the given address.
func (s *One_AccountsService) Delegate(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the delegate registration fee.
func (s *One_AccountsService) DelegateFee(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates/fee", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the balance for an account by the given address.
func (s *One_AccountsService) Balance(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getBalance", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the public key for an account by the given address.
func (s *One_AccountsService) PublicKey(ctx context.Context, address string, model interface{}) (interface{}, *http.Response, error) {
	query := &PublicKeyQuery{Address: address}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getPublicKey", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *One_AccountsService) Top(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/top", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

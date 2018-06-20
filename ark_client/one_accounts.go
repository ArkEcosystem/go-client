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

type PublicKeyResponse struct {
	Success   bool   `json:"success,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type Account struct {
	Address   string `json:"address,omitempty"`
	Balance   string `json:"balance,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type Accounts struct {
	Success  bool      `json:"success,omitempty"`
	Accounts []Account `json:"accounts,omitempty"`
}

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type One_AccountsService Service

// Get all accounts.
func (s *One_AccountsService) List(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getAllAccounts", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a account by the given address.
func (s *One_AccountsService) Get(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Count all accounts.
func (s *One_AccountsService) Count(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/count", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a delegate by the given address.
func (s *One_AccountsService) Delegate(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the delegate registration fee.
func (s *One_AccountsService) DelegateFee(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/delegates/fee", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the balance for an account by the given address.
func (s *One_AccountsService) Balance(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getBalance", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the public key for an account by the given address.
func (s *One_AccountsService) PublicKey(ctx context.Context, address string) (*PublicKeyResponse, *http.Response, error) {
	account := &PublicKeyResponse{}

	query := &PublicKeyQuery{Address: address}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/getPublicKey", query, &account)

	if err != nil {
		return nil, resp, err
	}

	return account, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *One_AccountsService) Top(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "accounts/top", nil, &accounts)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

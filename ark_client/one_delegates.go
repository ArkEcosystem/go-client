// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	"net/http"
)

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_DelegatesService Service

// Get all accounts.
func (s *One_DelegatesService) List(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a delegate by the given id.
func (s *One_DelegatesService) Get(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/get", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Count all delegates.
func (s *One_DelegatesService) Count(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/count", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the delegate registration fee.
func (s *One_DelegatesService) Fee(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/fee", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the forged totals by the given public key.
func (s *One_DelegatesService) ForgedByAccount(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/forging/getForgedByAccount", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Filter all delegates by the given criteria.
func (s *One_DelegatesService) Search(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/search", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all voters by the given public key.
func (s *One_DelegatesService) Voters(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/voters", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a list of the next forgers.
func (s *One_DelegatesService) NextForgers(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/getNextForgers", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

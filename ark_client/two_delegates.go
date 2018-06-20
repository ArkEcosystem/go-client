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

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type Two_DelegatesService Service

// Get all accounts.
func (s *Two_DelegatesService) List(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", "delegates", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get a block by the given id.
func (s *Two_DelegatesService) Get(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all blocks for the given delegate.
func (s *Two_DelegatesService) Blocks(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/blocks", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get all voters for the given delegate.
func (s *Two_DelegatesService) Voters(ctx context.Context, id int) (*Accounts, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/voters", id)

	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

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

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context, query *Pagination) (*Delegates, *http.Response, error) {
	var responseStruct *Delegates
	resp, err := s.client.SendRequest(ctx, 2, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *DelegatesService) Get(ctx context.Context, id int) (*Delegate, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v", id)

	var responseStruct *Delegate
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all blocks for the given delegate.
func (s *DelegatesService) Blocks(ctx context.Context, id int, query *Pagination) (*Blocks, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/blocks", id)

	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters for the given delegate.
func (s *DelegatesService) Voters(ctx context.Context, id int, query *Pagination) (*Wallets, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/voters", id)

	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

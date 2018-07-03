// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"github.com/ArkEcosystem/go-client/client"
	"net/http"
)

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type Two_DelegatesService Service

// Get all accounts.
func (s *Two_DelegatesService) List(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *Two_DelegatesService) Get(ctx context.Context, id int) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all blocks for the given delegate.
func (s *Two_DelegatesService) Blocks(ctx context.Context, id int, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/blocks", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters for the given delegate.
func (s *Two_DelegatesService) Voters(ctx context.Context, id int, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/voters", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
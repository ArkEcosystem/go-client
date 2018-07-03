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

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 2.
type Two_BlocksService Service

// Get all blocks.
func (s *Two_BlocksService) List(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "blocks", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *Two_BlocksService) Get(ctx context.Context, id int) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions by the given block.
func (s *Two_BlocksService) Transactions(ctx context.Context, id int, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v/transactions", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all blocks by the given criteria.
func (s *Two_BlocksService) Search(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "blocks/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

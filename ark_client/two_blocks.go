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

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 2.
type Two_BlocksService Service

// Get all blocks.
func (s *Two_BlocksService) List(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "blocks", query, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a block by the given id.
func (s *Two_BlocksService) Get(ctx context.Context, id int, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all transactions by the given block.
func (s *Two_BlocksService) Transactions(ctx context.Context, id int, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v/transactions", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, query, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Filter all blocks by the given criteria.
func (s *Two_BlocksService) Search(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "blocks/search", query, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

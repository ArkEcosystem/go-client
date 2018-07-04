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
type BlocksService Service

// Get all blocks.
func (s *BlocksService) List(ctx context.Context, query *Pagination) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", "blocks", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id int) (*Block, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v", id)

	var responseStruct *Block
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions by the given block.
func (s *BlocksService) Transactions(ctx context.Context, id int, query *Pagination) (*Transactions, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v/transactions", id)

	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all blocks by the given criteria.
func (s *BlocksService) Search(ctx context.Context, query *Pagination) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

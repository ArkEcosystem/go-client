// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

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
	resp, err := s.client.SendRequest(ctx, "GET", "blocks", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id int64) (*GetBlock, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v", id)

	var responseStruct *GetBlock
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the first block.
func (s *BlocksService) First(ctx context.Context) (*GetBlock, *http.Response, error) {
	var responseStruct *GetBlock
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/first", nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the last block.
func (s *BlocksService) Last(ctx context.Context) (*GetBlock, *http.Response, error) {
	var responseStruct *GetBlock
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/last", nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all transactions by the given block.
func (s *BlocksService) Transactions(ctx context.Context, id int64, query *Pagination) (*GetBlockTransactions, *http.Response, error) {
	uri := fmt.Sprintf("blocks/%v/transactions", id)

	var responseStruct *GetBlockTransactions
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

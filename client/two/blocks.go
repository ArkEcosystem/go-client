// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "fmt"
    "strings"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 2.
type BlocksService service

// Get all blocks.
func (s *BlocksService) List(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("GET", "blocks", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("blocks/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions by the given block.
func (s *BlocksService) Transactions(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("blocks/%v/transactions", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Filter all blocks by the given criteria.
func (s *BlocksService) Search(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("POST", "blocks/search", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

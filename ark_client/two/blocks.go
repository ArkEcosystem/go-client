// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "net/http"

    . "../types"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 2.
type BlocksService Service

// Get all blocks.
func (s *BlocksService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("blocks")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "blocks/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions by the given block.
func (s *BlocksService) Transactions(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "blocks/%v/transactions", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Filter all blocks by the given criteria.
func (s *BlocksService) Search(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("blocks/search")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

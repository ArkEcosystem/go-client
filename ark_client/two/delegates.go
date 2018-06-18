// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
)

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type DelegatesService service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "delegates", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *DelegatesService) Get(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("delegates/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all blocks for the given delegate.
func (s *DelegatesService) Blocks(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("delegates/%v/blocks", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all voters for the given delegate.
func (s *DelegatesService) Voters(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("delegates/%v/voters", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

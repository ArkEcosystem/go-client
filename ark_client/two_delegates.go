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
func (s *Two_DelegatesService) List(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(2, "GET", "delegates", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *Two_DelegatesService) Get(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("delegates/%v", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all blocks for the given delegate.
func (s *Two_DelegatesService) Blocks(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("delegates/%v/blocks", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all voters for the given delegate.
func (s *Two_DelegatesService) Voters(ctx context.Context, id int) (*http.Request, error) {
    uri := fmt.Sprintf("delegates/%v/voters", id)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

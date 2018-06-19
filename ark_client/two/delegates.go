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

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "delegates")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *DelegatesService) Get(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "delegates/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all blocks for the given delegate.
func (s *DelegatesService) Blocks(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "delegates/%v/blocks", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all voters for the given delegate.
func (s *DelegatesService) Voters(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "delegates/%v/voters", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

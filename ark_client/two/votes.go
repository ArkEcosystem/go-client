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

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type VotesService Service

// Get all votes.
func (s *VotesService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("votes")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a vote by the given id.
func (s *VotesService) Get(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "votes/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

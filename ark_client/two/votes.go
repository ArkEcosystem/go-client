// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "strings"
)

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type VotesService service

// Get all votes.
func (s *VotesService) List(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "votes", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a vote by the given id.
func (s *VotesService) Get(ctx context.Context) (*Response, error)
{
    uri := fmt.Sprintf("votes/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

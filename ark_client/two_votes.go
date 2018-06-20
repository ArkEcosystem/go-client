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

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type Two_VotesService Service

// Get all votes.
func (s *Two_VotesService) List(ctx context.Context) (*http.Response, error) {
    return s.client.SendRequest(ctx, 2, "GET", "votes", nil)
}


// Get a vote by the given id.
func (s *Two_VotesService) Get(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf("votes/%v", id)

    return s.client.SendRequest(ctx, 2, "GET", uri, nil)
}


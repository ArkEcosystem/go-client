// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"./structs/request"
	"context"
	"fmt"
	"net/http"
)

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type Two_VotesService Service

// Get all votes.
func (s *Two_VotesService) List(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "votes", query, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a vote by the given id.
func (s *Two_VotesService) Get(ctx context.Context, id int, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("votes/%v", id)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

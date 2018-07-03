// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"github.com/ArkEcosystem/go-client/client"
	"net/http"
)

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type Two_VotesService Service

// Get all votes.
func (s *Two_VotesService) List(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "votes", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a vote by the given id.
func (s *Two_VotesService) Get(ctx context.Context, id int) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("votes/%v", id)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

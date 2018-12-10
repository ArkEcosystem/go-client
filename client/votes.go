// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
)

// VotesService handles communication with the votes related
// methods of the Ark Core API - Version 2.
type VotesService Service

// Get all votes.
func (s *VotesService) List(ctx context.Context, query *Pagination) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "votes", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a vote by the given id.
func (s *VotesService) Get(ctx context.Context, id string) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("votes/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

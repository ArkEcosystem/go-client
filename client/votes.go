package client

import (
	"context"
	"fmt"
	"net/http"
)

type VotesService Service

// Get all votes.
func (s *VotesService) List(ctx context.Context, query *VotesQuery) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "votes", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a vote by the given id.
func (s *VotesService) Get(ctx context.Context, id string, query *VoteGetQuery) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("votes/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

package client

import (
	"context"
	"fmt"
	"net/http"
)

type RoundsService Service

// Get the forging validators of a round by the given id.
func (s *RoundsService) Validators(ctx context.Context, id string) (*Wallets, *http.Response, error) {
	uri := fmt.Sprintf("rounds/%v/validators", id)

	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all rounds.
func (s *RoundsService) All(ctx context.Context, query *Pagination) (*GetRounds, *http.Response, error) {
	uri := "rounds"

	var responseStruct *GetRounds
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a round by the given id.
func (s *RoundsService) Show(ctx context.Context, id string) (*GetRound, *http.Response, error) {
	uri := fmt.Sprintf("rounds/%v", id)

	var responseStruct *GetRound
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

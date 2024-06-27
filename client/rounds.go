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

// RoundsService handles communication with the rounds related
// methods of the Ark Core API - Version 2.
type RoundsService Service

// Get the forging delegates of a round by the given id.
func (s *RoundsService) Delegates(ctx context.Context, id int64) (*GetDelegates, *http.Response, error) {
	uri := fmt.Sprintf("rounds/%v/delegates", id)

	var responseStruct *GetDelegates
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all rounds.
func (s *RoundsService) All(ctx context.Context, query *Pagination) (*GetRounds, *http.Response, error) {
	uri := "rounds"

	var responseStruct *GetRounds
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a round by the given id.
func (s *RoundsService) Show(ctx context.Context, id int64) (*GetRound, *http.Response, error) {
	uri := fmt.Sprintf("rounds/%v", id)

	var responseStruct *GetRound
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
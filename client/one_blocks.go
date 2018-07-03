// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./requests/one"
	"./responses/one"
	"context"
	"net/http"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type One_BlocksService Service

// Get all accounts.
func (s *One_BlocksService) List(ctx context.Context, query *requests_one.BlocksQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *One_BlocksService) Get(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &requests_one.BlockIdQuery{Id: id}

	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain epoch.
func (s *One_BlocksService) Epoch(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getEpoch", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the transfer transaction fee.
func (s *One_BlocksService) Fee(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of transaction fees.
func (s *One_BlocksService) Fees(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFees", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain height.
func (s *One_BlocksService) Height(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getHeight", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain milestone.
func (s *One_BlocksService) Milestone(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getMilestone", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain nethash.
func (s *One_BlocksService) Nethash(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getNethash", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain reward.
func (s *One_BlocksService) Reward(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getReward", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain status.
func (s *One_BlocksService) Status(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getStatus", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain supply.
func (s *One_BlocksService) Supply(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getSupply", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

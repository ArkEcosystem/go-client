// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"./structs/request"
	"context"
	"net/http"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type One_BlocksService Service

// Get all accounts.
func (s *One_BlocksService) List(ctx context.Context, query *request.BlocksQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a block by the given id.
func (s *One_BlocksService) Get(ctx context.Context, id string, model interface{}) (interface{}, *http.Response, error) {
	query := &request.BlockIdQuery{Id: id}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/get", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain epoch.
func (s *One_BlocksService) Epoch(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getEpoch", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the transfer transaction fee.
func (s *One_BlocksService) Fee(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFee", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a list of transaction fees.
func (s *One_BlocksService) Fees(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFees", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain height.
func (s *One_BlocksService) Height(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getHeight", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain milestone.
func (s *One_BlocksService) Milestone(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getMilestone", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain nethash.
func (s *One_BlocksService) Nethash(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getNethash", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain reward.
func (s *One_BlocksService) Reward(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getReward", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain status.
func (s *One_BlocksService) Status(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getStatus", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the blockchain supply.
func (s *One_BlocksService) Supply(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getSupply", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

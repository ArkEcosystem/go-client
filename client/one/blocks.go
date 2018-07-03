// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"github.com/ArkEcosystem/go-client/client"
	"net/http"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type BlocksService Service

// Get all accounts.
func (s *BlocksService) List(ctx context.Context, query *one.BlocksQuery) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &one.BlockIdQuery{Id: id}

	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain epoch.
func (s *BlocksService) Epoch(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getEpoch", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the transfer transaction fee.
func (s *BlocksService) Fee(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of transaction fees.
func (s *BlocksService) Fees(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getFees", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain height.
func (s *BlocksService) Height(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getHeight", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain milestone.
func (s *BlocksService) Milestone(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getMilestone", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain nethash.
func (s *BlocksService) Nethash(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getNethash", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain reward.
func (s *BlocksService) Reward(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getReward", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain status.
func (s *BlocksService) Status(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getStatus", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain supply.
func (s *BlocksService) Supply(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "blocks/getSupply", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

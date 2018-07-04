// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"net/http"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type BlocksService Service

// Get all accounts.
func (s *BlocksService) List(ctx context.Context, query *BlocksQuery) (*Blocks, *http.Response, error) {
	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", "blocks", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context, id string) (*Block, *http.Response, error) {
	query := &BlockIdQuery{Id: id}

	var responseStruct *Block
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain epoch.
func (s *BlocksService) Epoch(ctx context.Context) (*BlocksEpoch, *http.Response, error) {
	var responseStruct *BlocksEpoch
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getEpoch", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the transfer transaction fee.
func (s *BlocksService) Fee(ctx context.Context) (*BlocksFee, *http.Response, error) {
	var responseStruct *BlocksFee
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getFee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of transaction fees.
func (s *BlocksService) Fees(ctx context.Context) (*BlocksFees, *http.Response, error) {
	var responseStruct *BlocksFees
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getFees", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain height.
func (s *BlocksService) Height(ctx context.Context) (*BlocksHeight, *http.Response, error) {
	var responseStruct *BlocksHeight
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getHeight", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain milest
func (s *BlocksService) Milestone(ctx context.Context) (*BlocksMilestone, *http.Response, error) {
	var responseStruct *BlocksMilestone
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getMilestone", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain nethash.
func (s *BlocksService) Nethash(ctx context.Context) (*BlocksNethash, *http.Response, error) {
	var responseStruct *BlocksNethash
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getNethash", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain reward.
func (s *BlocksService) Reward(ctx context.Context) (*BlocksReward, *http.Response, error) {
	var responseStruct *BlocksReward
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getReward", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain status.
func (s *BlocksService) Status(ctx context.Context) (*BlocksStatus, *http.Response, error) {
	var responseStruct *BlocksStatus
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getStatus", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the blockchain supply.
func (s *BlocksService) Supply(ctx context.Context) (*BlocksSupply, *http.Response, error) {
	var responseStruct *BlocksSupply
	resp, err := s.client.SendRequest(ctx, "GET", "blocks/getSupply", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

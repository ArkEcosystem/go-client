// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "net/http"

    . "../types"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type BlocksService Service

// Get all accounts.
func (s *BlocksService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/get")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain epoch.
func (s *BlocksService) Epoch(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getEpoch")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the transfer transaction fee.
func (s *BlocksService) Fee(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getFee")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a list of transaction fees.
func (s *BlocksService) Fees(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getFees")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain height.
func (s *BlocksService) Height(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getHeight")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain milestone.
func (s *BlocksService) Milestone(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getMilestone")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain nethash.
func (s *BlocksService) Nethash(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getNethash")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain reward.
func (s *BlocksService) Reward(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getReward")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain status.
func (s *BlocksService) Status(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getStatus")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain supply.
func (s *BlocksService) Supply(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/blocks/getSupply")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

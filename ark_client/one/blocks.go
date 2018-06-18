// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "fmt"
    "strings"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type BlocksService service

// Get all accounts.
func (s *BlocksService) List(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a block by the given id.
func (s *BlocksService) Get(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain epoch.
func (s *BlocksService) Epoch(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getEpoch", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the transfer transaction fee.
func (s *BlocksService) Fee(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getFee", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a list of transaction fees.
func (s *BlocksService) Fees(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getFees", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain height.
func (s *BlocksService) Height(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getHeight", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain milestone.
func (s *BlocksService) Milestone(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getMilestone", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain nethash.
func (s *BlocksService) Nethash(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getNethash", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain reward.
func (s *BlocksService) Reward(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getReward", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain status.
func (s *BlocksService) Status(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getStatus", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the blockchain supply.
func (s *BlocksService) Supply(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/blocks/getSupply", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// BlocksService handles communication with the blocks related
// methods of the Ark Core API - Version 1.
type One_BlocksService Service

// Get all accounts.
func (s *One_BlocksService) List(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a block by the given id.
func (s *One_BlocksService) Get(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/get", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain epoch.
func (s *One_BlocksService) Epoch(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getEpoch", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the transfer transaction fee.
func (s *One_BlocksService) Fee(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getFee", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a list of transaction fees.
func (s *One_BlocksService) Fees(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getFees", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain height.
func (s *One_BlocksService) Height(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getHeight", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain milestone.
func (s *One_BlocksService) Milestone(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getMilestone", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain nethash.
func (s *One_BlocksService) Nethash(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getNethash", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain reward.
func (s *One_BlocksService) Reward(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getReward", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain status.
func (s *One_BlocksService) Status(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getStatus", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the blockchain supply.
func (s *One_BlocksService) Supply(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "blocks/getSupply", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_DelegatesService Service

// Get all accounts.
func (s *One_DelegatesService) List(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a delegate by the given id.
func (s *One_DelegatesService) Get(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/get", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Count all delegates.
func (s *One_DelegatesService) Count(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/count", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the delegate registration fee.
func (s *One_DelegatesService) Fee(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/fee", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the forged totals by the given public key.
func (s *One_DelegatesService) ForgedByAccount(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/forging/getForgedByAccount", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Filter all delegates by the given criteria.
func (s *One_DelegatesService) Search(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/search", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get all voters by the given public key.
func (s *One_DelegatesService) Voters(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/voters", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a list of the next forgers.
func (s *One_DelegatesService) NextForgers(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "delegates/getNextForgers", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


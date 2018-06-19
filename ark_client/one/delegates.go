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

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a delegate by the given id.
func (s *DelegatesService) Get(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/get")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Count all delegates.
func (s *DelegatesService) Count(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/count")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the delegate registration fee.
func (s *DelegatesService) Fee(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/fee")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the forged totals by the given public key.
func (s *DelegatesService) ForgedByAccount(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/forging/getForgedByAccount")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Filter all delegates by the given criteria.
func (s *DelegatesService) Search(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/search")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all voters by the given public key.
func (s *DelegatesService) Voters(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/voters")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a list of the next forgers.
func (s *DelegatesService) NextForgers(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get("api/delegates/getNextForgers")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

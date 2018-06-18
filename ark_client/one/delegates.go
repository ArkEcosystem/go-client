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

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type DelegatesService service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a delegate by the given id.
func (s *DelegatesService) Get(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/get", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Count all delegates.
func (s *DelegatesService) Count(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/count", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the delegate registration fee.
func (s *DelegatesService) Fee(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/fee", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the forged totals by the given public key.
func (s *DelegatesService) ForgedByAccount(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/forging/getForgedByAccount", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Filter all delegates by the given criteria.
func (s *DelegatesService) Search(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/search", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get all voters by the given public key.
func (s *DelegatesService) Voters(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/voters", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a list of the next forgers.
func (s *DelegatesService) NextForgers(ctx context.Context) (*Response, error)
{
    req, err := s.client.NewRequest("GET", "api/delegates/getNextForgers", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

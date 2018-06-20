// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// NodeService handles communication with the node related
// methods of the Ark Core API - Version 2.
type Two_NodeService Service

// Get the node status.
func (s *Two_NodeService) Status(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(2, "GET", "node/status", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the node syncing status.
func (s *Two_NodeService) Syncing(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(2, "GET", "node/syncing", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get the node configuration.
func (s *Two_NodeService) Configuration(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(2, "GET", "node/configuration", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


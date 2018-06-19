// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "net/http"

    . "../types"
)

// NodeService handles communication with the node related
// methods of the Ark Core API - Version 2.
type NodeService Service

// Get the node status.
func (s *NodeService) Status(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "node/status")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get the node syncing status.
func (s *NodeService) Syncing(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "node/syncing")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get the node configuration.
func (s *NodeService) Configuration(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "node/configuration")

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

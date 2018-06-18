// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "strings"
)

// NodeService handles communication with the node related
// methods of the Ark Core API - Version 2.
type NodeService service

// Get the node status.
func (s *NodeService) Status(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "node/status", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get the node syncing status.
func (s *NodeService) Syncing(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "node/syncing", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get the node configuration.
func (s *NodeService) Configuration(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "node/configuration", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"net/http"
)

// NodeService handles communication with the node related
// methods of the Ark Core API - Version 2.
type NodeService Service

// Get the node status.
func (s *NodeService) Status(ctx context.Context) (*NodeStatus, *http.Response, error) {
	var responseStruct *NodeStatus
	resp, err := s.client.SendRequest(ctx, "GET", "node/status", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node syncing status.
func (s *NodeService) Syncing(ctx context.Context) (*NodeSyncing, *http.Response, error) {
	var responseStruct *NodeSyncing
	resp, err := s.client.SendRequest(ctx, "GET", "node/syncing", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node configuration.
func (s *NodeService) Configuration(ctx context.Context) (*NodeConfiguration, *http.Response, error) {
	var responseStruct *NodeConfiguration
	resp, err := s.client.SendRequest(ctx, "GET", "node/configuration", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

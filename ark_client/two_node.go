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
func (s *Two_NodeService) Status(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/status", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the node syncing status.
func (s *Two_NodeService) Syncing(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/syncing", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the node configuration.
func (s *Two_NodeService) Configuration(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/configuration", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

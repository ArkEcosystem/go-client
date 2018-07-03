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
type Two_NodeService Service

// Get the node status.
func (s *Two_NodeService) Status(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/status", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node syncing status.
func (s *Two_NodeService) Syncing(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/syncing", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node configuration.
func (s *Two_NodeService) Configuration(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "node/configuration", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./requests/one"
	"./responses/one"
	"context"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 1.
type One_PeersService Service

// Get all accounts.
func (s *One_PeersService) List(ctx context.Context, query *requests_one.GetPeersQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address and port.
func (s *One_PeersService) Get(ctx context.Context, query *requests_one.GetPeerQuery) (*http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers/get", query, &responseStruct)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get the node version of the given peer.
func (s *One_PeersService) Version(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers/version", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

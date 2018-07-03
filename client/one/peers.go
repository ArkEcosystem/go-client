// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 1.
type PeersService Service

// Get all accounts.
func (s *PeersService) List(ctx context.Context, query *GetPeersQuery) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address and port.
func (s *PeersService) Get(ctx context.Context, query *GetPeerQuery) (*http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers/get", query, &responseStruct)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get the node version of the given peer.
func (s *PeersService) Version(ctx context.Context) (*PublicKey, *http.Response, error) {
	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "peers/version", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

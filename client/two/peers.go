// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type PeersService Service

// Get all peers.
func (s *PeersService) List(ctx context.Context, query *Pagination) (*Peers, *http.Response, error) {
	var responseStruct *Peers
	resp, err := s.client.SendRequest(ctx, "GET", "peers", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address.
func (s *PeersService) Get(ctx context.Context, ip string) (*Peer, *http.Response, error) {
	uri := fmt.Sprintf("peers/%v", ip)

	var responseStruct *Peer
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

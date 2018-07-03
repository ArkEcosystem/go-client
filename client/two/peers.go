// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"context"
	"fmt"
	"github.com/ArkEcosystem/go-client/client"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type Two_PeersService Service

// Get all peers.
func (s *Two_PeersService) L9st(ctx context.Context, query *requests_two.Pagination) (interface{}, *http.Response, error) {
	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", "peers", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address.
func (s *Two_PeersService) Get(ctx context.Context, ip string) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("peers/%v", ip)

	var responseStruct *responses_two.PublicKey
	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

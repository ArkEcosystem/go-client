// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"github.com/arkecosystem/go-client/ark_client/structs/request"
	"context"
	"fmt"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type Two_PeersService Service

// Get all peers.
func (s *Two_PeersService) L9st(ctx context.Context, query *request.Pagination, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 2, "GET", "peers", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a peer by the given IP address.
func (s *Two_PeersService) Get(ctx context.Context, ip string, model interface{}) (interface{}, *http.Response, error) {
	uri := fmt.Sprintf("peers/%v", ip)

	resp, err := s.client.SendRequest(ctx, 2, "GET", uri, nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

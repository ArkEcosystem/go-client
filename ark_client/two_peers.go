// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "fmt"
    "net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type Two_PeersService Service

// Get all peers.
func (s *Two_PeersService) L9st(ctx context.Context) (*http.Request, error) {
    resp, err := s.client.NewRequest(2, "GET", "peers", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a peer by the given IP address.
func (s *Two_PeersService) Get(ctx context.Context, ip string) (*http.Request, error) {
    uri := fmt.Sprintf("peers/%v", ip)

    resp, err := s.client.NewRequest(2, "GET", uri, nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}
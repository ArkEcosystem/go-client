// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 1.
type One_PeersService Service

// Get all accounts.
func (s *One_PeersService) List(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "peers", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


// Get a peer by the given IP address and port.
// func (s *One_PeersService) Get(ctx context.Context) (*http.Response, error) {
//     req, err := s.client.NewRequest(1, "GET", "peers/get", compact("ip", "port"))

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get the node version of the given peer.
func (s *One_PeersService) Version(ctx context.Context) (*http.Response, error) {
    req, err := s.client.NewRequest(1, "GET", "peers/version", nil)

    if err != nil {
        return nil, err
    }

    resp, err := s.client.Do(ctx, req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}


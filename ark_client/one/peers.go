// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "net/http"

    . "../types"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 1.
type PeersService Service

// Get all accounts.
func (s *PeersService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "peers")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a peer by the given IP address and port.
// func (s *PeersService) Get(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "peers/get", compact("ip", "port"))

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get the node version of the given peer.
func (s *PeersService) Version(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "peers/version")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

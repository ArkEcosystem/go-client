// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "net/http"

    . "../types"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type PeersService Service

// Get all peers.
func (s *PeersService) L9st(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "peers")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get a peer by the given IP address.
func (s *PeersService) Get(ctx context.Context, ip string) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "peers/%v", ip)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

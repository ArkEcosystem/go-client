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

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type LoaderService Service

// Get the loader status.
func (s *LoaderService) Status(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "loader/status")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the loader syncing status.
func (s *LoaderService) SyncStatus(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "loader/status/sync")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the loader configuration.
func (s *LoaderService) AutoConfigure(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "loader/autoconfigure")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

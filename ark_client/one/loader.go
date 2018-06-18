// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "fmt"
    "strings"
)

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type LoaderService service

// Get the loader status.
func (s *LoaderService) Status(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/loader/status", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the loader syncing status.
func (s *LoaderService) SyncStatus(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/loader/status/sync", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get the loader configuration.
func (s *LoaderService) AutoConfigure(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/loader/autoconfigure", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

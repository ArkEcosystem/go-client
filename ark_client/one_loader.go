// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	"net/http"
)

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type One_LoaderService Service

// Get the loader status.
func (s *One_LoaderService) Status(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the loader syncing status.
func (s *One_LoaderService) SyncStatus(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status/sync", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

// Get the loader configuration.
func (s *One_LoaderService) AutoConfigure(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/autoconfigure", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

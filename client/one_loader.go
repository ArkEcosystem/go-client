// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"net/http"
)

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type One_LoaderService Service

// Get the loader status.
func (s *One_LoaderService) Status(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the loader syncing status.
func (s *One_LoaderService) SyncStatus(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status/sync", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the loader configuration.
func (s *One_LoaderService) AutoConfigure(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/autoconfigure", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"github.com/ArkEcosystem/go-client/client"
	"net/http"
)

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type LoaderService Service

// Get the loader status.
func (s *LoaderService) Status(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the loader syncing status.
func (s *LoaderService) SyncStatus(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/status/sync", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the loader configuration.
func (s *LoaderService) AutoConfigure(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "loader/autoconfigure", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

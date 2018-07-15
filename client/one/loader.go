// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"net/http"
)

// LoaderService handles communication with the loader related
// methods of the Ark Core API - Version 1.
type LoaderService Service

// Get the loader status.
func (s *LoaderService) Status(ctx context.Context) (*LoaderStatus, *http.Response, error) {
	var responseStruct *LoaderStatus
	resp, err := s.client.SendRequest(ctx, "GET", "loader/status", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the loader syncing status.
func (s *LoaderService) SyncStatus(ctx context.Context) (*LoaderSync, *http.Response, error) {
	var responseStruct *LoaderSync
	resp, err := s.client.SendRequest(ctx, "GET", "loader/status/sync", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the loader configuration.
func (s *LoaderService) AutoConfigure(ctx context.Context) (*LoaderAutoConfigure, *http.Response, error) {
	var responseStruct *LoaderAutoConfigure
	resp, err := s.client.SendRequest(ctx, "GET", "loader/autoconfigure", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

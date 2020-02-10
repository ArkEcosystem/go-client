// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
)

// BridgechainsService handles communication with the bridgechains related
// methods of the Ark Core API - Version 2.
type BridgechainsService Service

// Get all bridgechains.
func (s *BridgechainsService) List(ctx context.Context, query *Pagination) (*Bridgechains, *http.Response, error) {
	var responseStruct *Bridgechains
	resp, err := s.client.SendRequest(ctx, "GET", "bridgechains", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a bridgechains by the given id.
func (s *BridgechainsService) Get(ctx context.Context, id uint16) (*GetBridgechain, *http.Response, error) {
	uri := fmt.Sprintf("bridgechains/%v", id)

	var responseStruct *GetBridgechain
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all bridgechains by the given criteria.
func (s *BridgechainsService) Search(ctx context.Context, query *Pagination, body *BridgechainsSearchRequest) (*Bridgechains, *http.Response, error) {
	var responseStruct *Bridgechains
	resp, err := s.client.SendRequest(ctx, "POST", "bridgechains/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

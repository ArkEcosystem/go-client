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

// BusinessesService handles communication with the businesses related
// methods of the Ark Core API - Version 2.
type BusinessesService Service

// Get all businesses.
func (s *BusinessesService) List(ctx context.Context, query *Pagination) (*Businesses, *http.Response, error) {
	var responseStruct *Businesses
	resp, err := s.client.SendRequest(ctx, "GET", "businesses", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a business by the given id.
func (s *BusinessesService) Get(ctx context.Context, id uint16) (*GetBusiness, *http.Response, error) {
	uri := fmt.Sprintf("businesses/%v", id)

	var responseStruct *GetBusiness
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all bridgechains by the given business.
func (s *BusinessesService) Bridgechains(ctx context.Context, id uint16, query *Pagination) (*GetBusinessBridgechains, *http.Response, error) {
	uri := fmt.Sprintf("businesses/%v/bridgechains", id)

	var responseStruct *GetBusinessBridgechains
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all businesses by the given criteria.
func (s *BusinessesService) Search(ctx context.Context, query *Pagination, body *BusinessesSearchRequest) (*Businesses, *http.Response, error) {
	var responseStruct *Businesses
	resp, err := s.client.SendRequest(ctx, "POST", "businesses/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

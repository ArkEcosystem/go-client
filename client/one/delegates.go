// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"context"
	"net/http"
)

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context, query *one.GetDelegatesQuery) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a delegate by public key or username.
func (s *DelegatesService) Get(ctx context.Context, query *one.GetDelegateQuery) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Count all delegates.
func (s *DelegatesService) Count(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/count", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the delegate registration fee.
func (s *DelegatesService) Fee(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the forged totals by the given public key.
func (s *DelegatesService) ForgedByAccount(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &one.GeneratorPublicKeyQuery{GeneratorPublicKey: id}

	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/forging/getForgedByAccount", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all delegates by the given criteria.
// TODO: request query
func (s *DelegatesService) Search(ctx context.Context, query *one.DelegateSearchQuery) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters by the given public key.
func (s *DelegatesService) Voters(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &one.PublicKeyQuery{PublicKey: id}

	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/voters", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of the next forgers.
func (s *DelegatesService) NextForgers(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/getNextForgers", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

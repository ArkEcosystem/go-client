// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./requests/one"
	"./responses/one"
	"context"
	"net/http"
)

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_DelegatesService Service

// Get all accounts.
func (s *One_DelegatesService) List(ctx context.Context, query *requests_one.GetDelegatesQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a delegate by public key or username.
func (s *One_DelegatesService) Get(ctx context.Context, query *requests_one.GetDelegateQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Count all delegates.
func (s *One_DelegatesService) Count(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/count", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the delegate registration fee.
func (s *One_DelegatesService) Fee(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the forged totals by the given public key.
func (s *One_DelegatesService) ForgedByAccount(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &requests_one.GeneratorPublicKeyQuery{GeneratorPublicKey: id}

	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/forging/getForgedByAccount", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all delegates by the given criteria.
// TODO: request query
func (s *One_DelegatesService) Search(ctx context.Context, query *requests_one.DelegateSearchQuery) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters by the given public key.
func (s *One_DelegatesService) Voters(ctx context.Context, id string) (interface{}, *http.Response, error) {
	query := &requests_one.PublicKeyQuery{PublicKey: id}

	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/voters", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of the next forgers.
func (s *One_DelegatesService) NextForgers(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/getNextForgers", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

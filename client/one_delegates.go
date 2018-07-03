// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./structs/request"
	"context"
	"net/http"
)

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_DelegatesService Service

// Get all accounts.
func (s *One_DelegatesService) List(ctx context.Context, query *request.GetDelegatesQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a delegate by public key or username.
func (s *One_DelegatesService) Get(ctx context.Context, query *request.GetDelegateQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/get", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Count all delegates.
func (s *One_DelegatesService) Count(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/count", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the delegate registration fee.
func (s *One_DelegatesService) Fee(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/fee", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get the forged totals by the given public key.
func (s *One_DelegatesService) ForgedByAccount(ctx context.Context, id string, model interface{}) (interface{}, *http.Response, error) {
	query := &request.GeneratorPublicKeyQuery{GeneratorPublicKey: id}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/forging/getForgedByAccount", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Filter all delegates by the given criteria.
// TODO: request query
func (s *One_DelegatesService) Search(ctx context.Context, query *request.DelegateSearchQuery, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/search", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get all voters by the given public key.
func (s *One_DelegatesService) Voters(ctx context.Context, id string, model interface{}) (interface{}, *http.Response, error) {
	query := &request.PublicKeyQuery{PublicKey: id}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/voters", query, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

// Get a list of the next forgers.
func (s *One_DelegatesService) NextForgers(ctx context.Context, model interface{}) (interface{}, *http.Response, error) {
	resp, err := s.client.SendRequest(ctx, 1, "GET", "delegates/getNextForgers", nil, &model)

	if err != nil {
		return nil, resp, err
	}

	return model, resp, err
}

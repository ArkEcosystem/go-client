package client

import (
	"context"
	"fmt"
	"net/http"
)

// ValidatorsService handles communication with the validators related
// methods of the Ark Core API - Version 2.
type ValidatorsService Service

// Get all accounts.
func (s *ValidatorsService) List(ctx context.Context, query *Pagination) (*Validators, *http.Response, error) {
	var responseStruct *Validators
	resp, err := s.client.SendRequest(ctx, "GET", "validators", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a validator by the given ID. (address, publicKey and username are valid)
func (s *ValidatorsService) Get(ctx context.Context, id string) (*GetValidator, *http.Response, error) {
	uri := fmt.Sprintf("validators/%v", id)

	var responseStruct *GetValidator
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all blocks for the given validator.
func (s *ValidatorsService) Blocks(ctx context.Context, id string, query *Pagination) (*GetValidatorBlocks, *http.Response, error) {
	uri := fmt.Sprintf("validators/%v/blocks", id)

	var responseStruct *GetValidatorBlocks
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters for the given validator.
func (s *ValidatorsService) Voters(ctx context.Context, id string, query *Pagination) (*GetValidatorVoters, *http.Response, error) {
	uri := fmt.Sprintf("validators/%v/voters", id)

	var responseStruct *GetValidatorVoters
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

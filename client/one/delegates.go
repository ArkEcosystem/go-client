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

// DelegatesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context, query *GetDelegatesQuery) (*AccountDelegates, *http.Response, error) {
	var responseStruct *AccountDelegates
	resp, err := s.client.SendRequest(ctx, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a delegate by public key or username.
func (s *DelegatesService) Get(ctx context.Context, query *GetDelegateQuery) (*GetDelegate, *http.Response, error) {
	var responseStruct *GetDelegate
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Count all delegates.
func (s *DelegatesService) Count(ctx context.Context) (*DelegatesCount, *http.Response, error) {
	var responseStruct *DelegatesCount
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/count", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the delegate registration fee.
func (s *DelegatesService) Fee(ctx context.Context) (*DelegateFee, *http.Response, error) {
	var responseStruct *DelegateFee
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the forged totals by the given public key.
func (s *DelegatesService) ForgedByAccount(ctx context.Context, id string) (*ForgedByDelegate, *http.Response, error) {
	query := &GeneratorPublicKeyQuery{GeneratorPublicKey: id}

	var responseStruct *ForgedByDelegate
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/forging/getForgedByAccount", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all delegates by the given criteria.
func (s *DelegatesService) Search(ctx context.Context, query *DelegateSearchQuery) (*AccountDelegates, *http.Response, error) {
	var responseStruct *AccountDelegates
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/search", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters by the given public key.
func (s *DelegatesService) Voters(ctx context.Context, id string) (*Accounts, *http.Response, error) {
	query := &PublicKeyQuery{PublicKey: id}

	var responseStruct *Accounts
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/voters", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a list of the next forgers.
func (s *DelegatesService) NextForgers(ctx context.Context) (*NextForger, *http.Response, error) {
	var responseStruct *NextForger
	resp, err := s.client.SendRequest(ctx, "GET", "delegates/getNextForgers", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

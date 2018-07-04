// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

import (
	"context"
	"fmt"
	"net/http"
)

// DelegatesService handles communication with the delegates related
// methods of the Ark Core API - Version 2.
type DelegatesService Service

// Get all accounts.
func (s *DelegatesService) List(ctx context.Context, query *Pagination) (*Delegates, *http.Response, error) {
	var responseStruct *Delegates
	resp, err := s.client.SendRequest(ctx, "GET", "delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a block by the given id.
func (s *DelegatesService) Get(ctx context.Context, id int) (*Delegate, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v", id)

	var responseStruct *Delegate
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all blocks for the given delegate.
func (s *DelegatesService) Blocks(ctx context.Context, id int, query *Pagination) (*Blocks, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/blocks", id)

	var responseStruct *Blocks
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all voters for the given delegate.
func (s *DelegatesService) Voters(ctx context.Context, id int, query *Pagination) (*Wallets, *http.Response, error) {
	uri := fmt.Sprintf("delegates/%v/voters", id)

	var responseStruct *Wallets
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

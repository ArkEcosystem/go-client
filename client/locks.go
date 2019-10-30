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

// LocksService handles communication with the locks related
// methods of the Ark Core API - Version 2.
type LocksService Service

// Get all locks.
func (s *LocksService) List(ctx context.Context, query *Pagination) (*Locks, *http.Response, error) {
	var responseStruct *Locks
	resp, err := s.client.SendRequest(ctx, "GET", "locks", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a lock by the given id.
func (s *LocksService) Get(ctx context.Context, id string) (*GetLock, *http.Response, error) {
	uri := fmt.Sprintf("locks/%v", id)

	var responseStruct *GetLock
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Filter all locks by the given criteria.
func (s *LocksService) Search(ctx context.Context, query *Pagination, body *LocksSearchRequest) (*Locks, *http.Response, error) {
	var responseStruct *Locks
	resp, err := s.client.SendRequest(ctx, "POST", "locks/search", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Retrieve transactions by the given lock ids.
func (s *LocksService) Unlocked(ctx context.Context, query *Pagination, body *LocksUnlockedRequest) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "POST", "locks/unlocked", query, body, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

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

// EntitiesService
type EntitiesService Service

// Get all entities.
func (s *EntitiesService) List(ctx context.Context, query *Pagination) (*Entities, *http.Response, error) {
	var responseStruct *Entities
	resp, err := s.client.SendRequest(ctx, "GET", "entities", query, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an entity by the given id.
func (s *EntitiesService) Get(ctx context.Context, id string) (*GetEntity, *http.Response, error) {
	uri := fmt.Sprintf("entities/%v", id)

	var responseStruct *GetEntity
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

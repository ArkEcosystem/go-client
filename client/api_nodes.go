// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"net/http"
)

// ApiNodesService handles communication with the API nodes related
// methods of the Ark Core API.
type ApiNodesService Service

// Get all available nodes serving APIs.
func (s *ApiNodesService) All(ctx context.Context, query *Pagination) (*ApiNodesResponse, *http.Response, error) {
	var responseStruct *ApiNodesResponse
	resp, err := s.client.SendRequest(ctx, "GET", "api-nodes", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}


package client

import (
	"context"
	"net/http"
)

type ApiNodesService Service

// Get all available nodes serving APIs.
func (s *ApiNodesService) All(ctx context.Context, query *ApiNodesQuery) (*ApiNodesResponse, *http.Response, error) {
	var responseStruct *ApiNodesResponse
	resp, err := s.client.SendRequest(ctx, "GET", "api-nodes", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

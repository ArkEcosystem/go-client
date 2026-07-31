package client

import (
	"context"
	"net/http"
)

type NodeService Service

// Get the node status.
func (s *NodeService) Status(ctx context.Context) (*GetNodeStatus, *http.Response, error) {
	var responseStruct *GetNodeStatus
	resp, err := s.client.SendRequest(ctx, "GET", "node/status", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node syncing status.
func (s *NodeService) Syncing(ctx context.Context) (*GetNodeSyncing, *http.Response, error) {
	var responseStruct *GetNodeSyncing
	resp, err := s.client.SendRequest(ctx, "GET", "node/syncing", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node configuration.
func (s *NodeService) Configuration(ctx context.Context) (*GetNodeConfiguration, *http.Response, error) {
	var responseStruct *GetNodeConfiguration
	resp, err := s.client.SendRequest(ctx, "GET", "node/configuration", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *NodeService) Crypto(ctx context.Context) (*GetNodeCrypto, *http.Response, error) {
	var responseStruct *GetNodeCrypto
	resp, err := s.client.SendRequest(ctx, "GET", "node/configuration/crypto", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *NodeService) Fees(ctx context.Context, days int) (*GetNodeFees, *http.Response, error) {
	var responseStruct *GetNodeFees
	resp, err := s.client.SendRequest(ctx, "GET", "node/fees", FeesRequest{days}, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

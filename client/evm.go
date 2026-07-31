package client

import (
	"context"
	"net/http"
)

type EVMService Service

type EvmRequest struct {
	Jsonrpc string        `json:"jsonrpc,omitempty"`
	Method  string        `json:"method,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
	Id      *int          `json:"id"`
}

func (s *EVMService) Call(ctx context.Context, req *EvmRequest) (map[string]interface{}, *http.Response, error) {
	if req.Jsonrpc == "" {
		req.Jsonrpc = "2.0"
	}

	var responseStruct map[string]interface{}
	resp, err := s.client.SendRequest(ctx, "POST", "", nil, req, &responseStruct, "evm")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

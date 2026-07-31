package client

import (
	"context"
	"fmt"
	"net/http"
)

type ContractsService Service

func (s *ContractsService) All(ctx context.Context) (*ContractsResponse, *http.Response, error) {
	var responseStruct *ContractsResponse
	resp, err := s.client.SendRequest(ctx, "GET", "contracts", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *ContractsService) Abi(ctx context.Context, name string, implementation string) (*ContractAbiResponse, *http.Response, error) {
	uri := fmt.Sprintf("contracts/%v/%v/abi", name, implementation)

	var responseStruct *ContractAbiResponse
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

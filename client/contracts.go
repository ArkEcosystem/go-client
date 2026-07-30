package client

import (
	"context"
	"fmt"
	"net/http"
)

// ContractsService handles communication with the EVM contracts related
// methods of the Ark Core API.
type ContractsService Service

// Get all deployed contracts, keyed by name.
func (s *ContractsService) All(ctx context.Context) (*ContractsResponse, *http.Response, error) {
	var responseStruct *ContractsResponse
	resp, err := s.client.SendRequest(ctx, "GET", "contracts", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the ABI for the given contract name and implementation address.
func (s *ContractsService) Abi(ctx context.Context, name string, implementation string) (*ContractAbiResponse, *http.Response, error) {
	uri := fmt.Sprintf("contracts/%v/%v/abi", name, implementation)

	var responseStruct *ContractAbiResponse
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

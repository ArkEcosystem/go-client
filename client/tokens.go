package client

import (
	"context"
	"fmt"
	"net/http"
)

type TokensService Service

func (s *TokensService) All(ctx context.Context, query *TokensQuery) (*Tokens, *http.Response, error) {
	var responseStruct *Tokens
	resp, err := s.client.SendRequest(ctx, "GET", "tokens", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) Transfers(ctx context.Context, query *TokenTransfersQuery) (*TokenActions, *http.Response, error) {
	var responseStruct *TokenActions
	resp, err := s.client.SendRequest(ctx, "GET", "tokens/transfers", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) Approvals(ctx context.Context, query *TokenApprovalsQuery) (*TokenActions, *http.Response, error) {
	var responseStruct *TokenActions
	resp, err := s.client.SendRequest(ctx, "GET", "tokens/approvals", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) Whitelist(ctx context.Context, query *Pagination) (*TokenWhitelist, *http.Response, error) {
	var responseStruct *TokenWhitelist
	resp, err := s.client.SendRequest(ctx, "GET", "tokens/whitelist", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) Get(ctx context.Context, contractAddress string) (*GetToken, *http.Response, error) {
	uri := fmt.Sprintf("tokens/%v", contractAddress)

	var responseStruct *GetToken
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) TransfersFor(ctx context.Context, contractAddress string, query *TokenLookupQuery) (*TokenActionsResults, *http.Response, error) {
	uri := fmt.Sprintf("tokens/%v/transfers", contractAddress)

	var responseStruct *TokenActionsResults
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) ApprovalsFor(ctx context.Context, contractAddress string, query *TokenLookupQuery) (*TokenActionsResults, *http.Response, error) {
	uri := fmt.Sprintf("tokens/%v/approvals", contractAddress)

	var responseStruct *TokenActionsResults
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TokensService) HoldersFor(ctx context.Context, contractAddress string) (*TokenHolders, *http.Response, error) {
	uri := fmt.Sprintf("tokens/%v/holders", contractAddress)

	var responseStruct *TokenHolders
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

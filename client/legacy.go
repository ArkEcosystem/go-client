package client

import (
	"context"
	"fmt"
	"net/http"
)

type LegacyService Service

func (s *LegacyService) ColdWallets(ctx context.Context, query *Pagination) (*LegacyColdWallets, *http.Response, error) {
	var responseStruct *LegacyColdWallets
	resp, err := s.client.SendRequest(ctx, "GET", "legacy/cold-wallets", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *LegacyService) ColdWallet(ctx context.Context, address string) (*GetLegacyColdWallet, *http.Response, error) {
	uri := fmt.Sprintf("legacy/cold-wallets/%v", address)

	var responseStruct *GetLegacyColdWallet
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

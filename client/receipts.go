package client

import (
	"context"
	"fmt"
	"net/http"
)

type ReceiptsService Service

func (s *ReceiptsService) All(ctx context.Context, query *Pagination) (*Receipts, *http.Response, error) {
	var responseStruct *Receipts
	resp, err := s.client.SendRequest(ctx, "GET", "receipts", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *ReceiptsService) Get(ctx context.Context, transactionHash string) (*GetReceipt, *http.Response, error) {
	uri := fmt.Sprintf("receipts/%v", transactionHash)

	var responseStruct *GetReceipt
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *ReceiptsService) Contracts(ctx context.Context, query *Pagination) (*Receipts, *http.Response, error) {
	var responseStruct *Receipts
	resp, err := s.client.SendRequest(ctx, "GET", "receipts/contracts", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

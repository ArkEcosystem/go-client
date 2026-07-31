package client

import (
	"context"
	"fmt"
	"net/http"
)

type TransactionsService Service

// Get all transactions.
func (s *TransactionsService) List(ctx context.Context, query *TransactionsQuery) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Create a new transaction.
func (s *TransactionsService) Create(ctx context.Context, body *CreateTransactionRequest) (*CreateTransaction, *http.Response, error) {
	var responseStruct *CreateTransaction
	resp, err := s.client.SendRequest(ctx, "POST", "transactions", nil, body, &responseStruct, "transactions")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id string, query *TransactionGetQuery) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context, query *UnconfirmedTransactionsQuery) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/unconfirmed", query, nil, &responseStruct, "transactions")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id string) (*GetTransaction, *http.Response, error) {
	uri := fmt.Sprintf("transactions/unconfirmed/%v", id)

	var responseStruct *GetTransaction
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "transactions")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

func (s *TransactionsService) Configuration(ctx context.Context) (*GetTransactionConfiguration, *http.Response, error) {
	var responseStruct *GetTransactionConfiguration
	resp, err := s.client.SendRequest(ctx, "GET", "configuration", nil, nil, &responseStruct, "transactions")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

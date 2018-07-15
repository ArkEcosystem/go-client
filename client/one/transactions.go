// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"net/http"
)

// TransactionsService handles communication with the transactions
// methods of the Ark Core API - Version 1.
type TransactionsService Service

// Get all accounts.
func (s *TransactionsService) List(ctx context.Context, query *GetTransactionsQuery) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a transaction by the given id.
func (s *TransactionsService) Get(ctx context.Context, id string) (*Transaction, *http.Response, error) {
	query := &TransactionIdQuery{Id: id}

	var responseStruct *Transaction
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all unconfirmed transactions.
func (s *TransactionsService) ListUnconfirmed(ctx context.Context, query *GetUnconfirmedTransactionsQuery) (*Transactions, *http.Response, error) {
	var responseStruct *Transactions
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/unconfirmed", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get an unconfirmed transaction by the given id.
func (s *TransactionsService) GetUnconfirmed(ctx context.Context, id string) (*Transaction, *http.Response, error) {
	query := &TransactionIdQuery{Id: id}

	var responseStruct *Transaction
	resp, err := s.client.SendRequest(ctx, "GET", "transactions/unconfirmed/get", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

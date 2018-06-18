// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "fmt"
    "strings"
)

// WalletsService handles communication with the wallets related
// methods of the Ark Core API - Version 2.
type WalletsService service

// Get all wallets.
func (s *WalletsService) List(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("GET", "wallets", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all wallets sorted by balance in descending order.
func (s *WalletsService) Top(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("GET", "wallets/top", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a wallet by the given id.
func (s *WalletsService) Get(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("wallets/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions for the given wallet.
func (s *WalletsService) Transactions(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("wallets/%v/transactions", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions sent by the given wallet.
func (s *WalletsService) SentTransactions(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("wallets/%v/transactions/sent", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all transactions received by the given wallet.
func (s *WalletsService) ReceivedTransaction(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("wallets/%v/transactions/received", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get all votes by the given wallet.
func (s *WalletsService) Votes(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("wallets/%v/votes", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Filter all wallets by the given criteria.
func (s *WalletsService) Search(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("POST", "wallets/search", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

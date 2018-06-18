// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "strings"
)

// WebhooksService handles communication with the webhooks related
// methods of the Ark Core API - Version 2.
type WebhooksService service

// Get all webhooks.
func (s *WebhooksService) List(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "webhooks", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Create a new webhook.
func (s *WebhooksService) Create(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("POST", "webhooks", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get the webhook by the given id.
func (s *WebhooksService) Show(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("webhooks/%v", id)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Update the webhook by the given id.
func (s *WebhooksService) Update(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("webhooks/%v", id)

    req, err := s.client.NewRequest("PUT", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Delete the webhook by the given id.
func (s *WebhooksService) Delete(ctx context.Context) (*Response, error) {
    uri := fmt.Sprintf("webhooks/%v", id)

    req, err := s.client.NewRequest("DELETE", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
    "context"
    "fmt"
    "net/http"

    . "../types"
)

// WebhooksService handles communication with the webhooks related
// methods of the Ark Core API - Version 2.
type WebhooksService Service

// Get all webhooks.
func (s *WebhooksService) List(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "webhooks")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Create a new webhook.
// func (s *WebhooksService) Create(ctx context.Context) (*http.Response, error) {
//     resp, err := s.Client.Client.Post(s.Client.BaseURL.String() + "webhooks")

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// Get the webhook by the given id.
func (s *WebhooksService) Show(ctx context.Context, id int) (*http.Response, error) {
    uri := fmt.Sprintf(s.Client.BaseURL.String() + "webhooks/%v", id)

    resp, err := s.Client.Client.Get(uri)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Update the webhook by the given id.
// func (s *WebhooksService) Update(ctx context.Context, id int) (*http.Response, error) {
//     uri := fmt.Sprintf(s.Client.BaseURL.String() + "webhooks/%v", id)

//     resp, err := s.Client.Client.Put(uri)

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

// // Delete the webhook by the given id.
// func (s *WebhooksService) Delete(ctx context.Context, id int) (*http.Response, error) {
//     uri := fmt.Sprintf(s.Client.BaseURL.String() + "webhooks/%v", id)

//     resp, err := s.Client.Client.Delete(uri)

//     if err != nil {
//         return nil, err
//     }

//     return resp, nil
// }

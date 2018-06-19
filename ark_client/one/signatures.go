// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "net/http"

    . "../types"
)

// SignaturesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type SignaturesService Service

// Get the second signature registration fee.
func (s *SignaturesService) Fee(ctx context.Context) (*http.Response, error) {
    resp, err := s.Client.Client.Get(s.Client.BaseURL.String() + "api/signatures/fee")

    if err != nil {
        return nil, err
    }

    return resp, nil
}

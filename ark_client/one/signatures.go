// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
    "context"
    "fmt"
    "strings"
)

// SignaturesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type SignaturesService service

// Get the second signature registration fee.
func (s *SignaturesService) Fee(ctx context.Context) (*Response, error) {
    req, err := s.client.NewRequest("GET", "api/signatures/fee", nil)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

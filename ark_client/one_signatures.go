// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
	"context"
	"net/http"
)

// SignaturesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_SignaturesService Service

// Get the second signature registration fee.
func (s *One_SignaturesService) Fee(ctx context.Context) (*Accounts, *http.Response, error) {
	accounts := &Accounts{}

	resp, err := s.client.SendRequest(ctx, 1, "GET", "signatures/fee", nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accounts, resp, err
}

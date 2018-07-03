// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package client

import (
	"./responses/one"
	"context"
	"net/http"
)

// SignaturesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type One_SignaturesService Service

// Get the second signature registration fee.
func (s *One_SignaturesService) Fee(ctx context.Context) (interface{}, *http.Response, error) {
	var responseStruct *responses_one.PublicKey
	resp, err := s.client.SendRequest(ctx, 1, "GET", "signatures/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

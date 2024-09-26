// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"net/http"
)

type BlockchainService Service

// Get blockchain information.
func (s *BlockchainService) Info(ctx context.Context) (*BlockchainInfo, *http.Response, error) {
	var responseStruct *BlockchainInfo
	resp, err := s.client.SendRequest(ctx, "GET", "blockchain", nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

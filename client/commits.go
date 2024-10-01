// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"context"
	"fmt"
	"net/http"
)

type CommitsService Service

// GetCommit takes a commit height and returns the commit details.
func (s *CommitsService) GetCommit(ctx context.Context, height int) (*CommitResponse, *http.Response, error) {
	uri := fmt.Sprintf("commits/%d", height)

	var responseStruct *CommitResponse
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

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

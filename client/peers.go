package client

import (
	"context"
	"fmt"
	"net/http"
)

type PeersService Service

// Get all peers.
func (s *PeersService) List(ctx context.Context, query *PeersQuery) (*Peers, *http.Response, error) {
	var responseStruct *Peers
	resp, err := s.client.SendRequest(ctx, "GET", "peers", query, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address.
func (s *PeersService) Get(ctx context.Context, ip string) (*GetPeer, *http.Response, error) {
	uri := fmt.Sprintf("peers/%v", ip)

	var responseStruct *GetPeer
	resp, err := s.client.SendRequest(ctx, "GET", uri, nil, nil, &responseStruct, "api")

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"net/http"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 1.
type PeersService Service

// Get all accounts.
func (s *PeersService) List(ctx context.Context, query *GetPeersQuery) (*Peers, *http.Response, error) {
	var responseStruct *Peers
	resp, err := s.client.SendRequest(ctx, "GET", "peers", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a peer by the given IP address and port.
func (s *PeersService) Get(ctx context.Context, query *GetPeerQuery) (*GetPeer, *http.Response, error) {
	var responseStruct *GetPeer
	resp, err := s.client.SendRequest(ctx, "GET", "peers/get", query, &responseStruct)

	if err != nil {
		return nil, nil, err
	}

	return responseStruct, resp, nil
}

// Get the node version of the given peer.
func (s *PeersService) Version(ctx context.Context) (*PeersVersion, *http.Response, error) {
	var responseStruct *PeersVersion
	resp, err := s.client.SendRequest(ctx, "GET", "peers/version", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

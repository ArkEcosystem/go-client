// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all accounts.
func TestPeersService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "peers": [
			  	{
			      "ip": "1.2.3.4",
			      "port": 4002,
			      "version": "2.0.0",
			      "errors": 0,
			      "os": "dummy",
			      "height": 100,
			      "status": "OK",
			      "delay": 10
			    }
			  ]
			}`)
	})

	query := &GetPeersQuery{Limit: 1}
	responseStruct, response, err := client.Peers.List(context.Background(), query)
	testGeneralError(t, "Peers.List", err)
	testResponseUrl(t, "Peers.List", response, "/api/peers")
	testResponseStruct(t, "Peers.List", responseStruct, &Peers{
		Success: true,
		Peers: []Peer{{
			Ip:      "1.2.3.4",
			Port:    4002,
			Version: "2.0.0",
			Errors:  0,
			Os:      "dummy",
			Height:  100,
			Status:  "OK",
			Delay:   10,
		}},
	})
}

// Get a peer by the given IP address and port.
func TestPeersService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers/get", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "peer": {
			    "ip": "1.2.3.4",
			    "port": 4002,
			    "version": "2.0.0",
			    "errors": 0,
			    "os": "dummy",
			    "height": 100,
			    "status": "OK",
			    "delay": 10
			  }
			}`)
	})

	query := &GetPeerQuery{Ip: "1.2.3.4", Port: 4002}
	responseStruct, response, err := client.Peers.Get(context.Background(), query)
	testGeneralError(t, "Peers.Get", err)
	testResponseUrl(t, "Peers.Get", response, "/api/peers/get")
	testResponseStruct(t, "Peers.Get", responseStruct, &GetPeer{
		Success: true,
		Peer: Peer{
			Ip:      "1.2.3.4",
			Port:    4002,
			Version: "2.0.0",
			Errors:  0,
			Os:      "dummy",
			Height:  100,
			Status:  "OK",
			Delay:   10,
		},
	})
}

// Get the node version of the given peer.
func TestPeersService_Version(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers/version", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "version": "2.0.0",
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Peers.Version(context.Background())
	testGeneralError(t, "Peers.Version", err)
	testResponseUrl(t, "Peers.Version", response, "/api/peers/version")
	testResponseStruct(t, "Peers.Version", responseStruct, &PeersVersion{
		Success: true,
		Version: "2.0.0",
	})
}

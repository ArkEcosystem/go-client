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
	"testing"
)

// Get all peers.
func TestPeersService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/peers?page=1&limit=1",
			    "first": "/api/peers?page=1&limit=1",
			    "last": "/api/peers?page=1&limit=1"
			  },
			  "data": [
			    {
			      "ip": "1.2.3.4",
			      "port": 4002,
			      "version": "2.0.0",
			      "status": "OK",
			      "os": "dummy",
			      "latency": 10
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Peers.List(context.Background(), query)
	testGeneralError(t, "Peers.List", err)
	testResponseUrl(t, "Peers.List", response, "/api/peers")
	testResponseStruct(t, "Peers.List", responseStruct, &Peers{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/peers?page=1&limit=1",
			First:      "/api/peers?page=1&limit=1",
			Last:       "/api/peers?page=1&limit=1",
		},
		Data: []Peer{{
			Ip:      "1.2.3.4",
			Port:    4002,
			Version: "2.0.0",
			Status:  "OK",
			Os:      "dummy",
			Latency: 10,
		}},
	})
}

// Get a peer by the given IP address.
func TestPeersService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers/1.2.3.4", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/peers?page=1&limit=1",
			    "first": "/api/peers?page=1&limit=1",
			    "last": "/api/peers?page=1&limit=1"
			  },
			  "data": {
			    "ip": "1.2.3.4",
			    "port": 4002,
			    "version": "2.0.0",
			    "status": "OK",
			    "os": "dummy",
			    "latency": 10
			  }
			}`)
	})

	responseStruct, response, err := client.Peers.Get(context.Background(), "1.2.3.4")
	testGeneralError(t, "Peers.Get", err)
	testResponseUrl(t, "Peers.Get", response, "/api/peers/1.2.3.4")
	testResponseStruct(t, "Peers.Get", responseStruct, &GetPeer{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/peers?page=1&limit=1",
			First:      "/api/peers?page=1&limit=1",
			Last:       "/api/peers?page=1&limit=1",
		},
		Data: Peer{
			Ip:      "1.2.3.4",
			Port:    4002,
			Version: "2.0.0",
			Status:  "OK",
			Os:      "dummy",
			Latency: 10,
		},
	})
}

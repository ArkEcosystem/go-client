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
			    "totalCountIsEstimate": false,
			    "count": 1,
			    "first": "/peers?limit=1&page=1",
			    "last": "/peers?limit=1&page=12",
			    "next": "/peers?limit=1&page=2",
			    "pageCount": 12,
			    "previous": null,
			    "self": "/peers?limit=1&page=1",
			    "totalCount": 12
			  },
			  "data": [
			    {
			      "blockNumber": 23189863,
			      "ip": "128.140.81.247",
			      "latency": 3,
			      "plugins": {
			        "@mainsail/webhooks": {
			          "port": 4004,
			          "enabled": false,
			          "estimateTotalCount": false
			        },
			        "@mainsail/api-development": {
			          "port": 4006,
			          "enabled": true,
			          "estimateTotalCount": false
			        }
			      },
			      "port": 4000,
			      "ports": {
			        "@mainsail/webhooks": -1,
			        "@mainsail/api-development": 4006
			      },
			      "version": "0.0.1-evm.53"
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
			TotalCountIsEstimate: false,
			Count:                1,
			First:                "/peers?limit=1&page=1",
			Last:                 "/peers?limit=1&page=12",
			Next:                 strPtr("/peers?limit=1&page=2"),
			PageCount:            12,
			Previous:             nil,
			Self:                 "/peers?limit=1&page=1",
			TotalCount:           12,
		},
		Data: []Peer{{
			BlockNumber: 23189863,
			Ip:          "128.140.81.247",
			Latency:     3,
			Plugins: map[string]PeerPlugin{
				"@mainsail/webhooks": {
					Port:               4004,
					Enabled:            false,
					EstimateTotalCount: false,
				},
				"@mainsail/api-development": {
					Port:               4006,
					Enabled:            true,
					EstimateTotalCount: false,
				},
			},
			Port: 4000,
			Ports: PeerPorts{
				"@mainsail/webhooks":        -1,
				"@mainsail/api-development": 4006,
			},
			Version: "0.0.1-evm.53",
		}},
	})
}

// Get a peer by the given IP address.
func TestPeersService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/peers/128.140.81.247", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "blockNumber": 23189864,
			    "ip": "128.140.81.247",
			    "latency": 67,
			    "plugins": {
			      "@mainsail/webhooks": {
			        "port": 4004,
			        "enabled": false,
			        "estimateTotalCount": false
			      },
			      "@mainsail/api-development": {
			        "port": 4006,
			        "enabled": true,
			        "estimateTotalCount": false
			      }
			    },
			    "port": 4000,
			    "ports": {
			      "@mainsail/webhooks": -1,
			      "@mainsail/api-development": 4006
			    },
			    "version": "0.0.1-evm.53"
			  }
			}`)
	})

	responseStruct, response, err := client.Peers.Get(context.Background(), "128.140.81.247")
	testGeneralError(t, "Peers.Get", err)
	testResponseUrl(t, "Peers.Get", response, "/api/peers/128.140.81.247")
	testResponseStruct(t, "Peers.Get", responseStruct, &GetPeer{
		Data: Peer{
			BlockNumber: 23189864,
			Ip:          "128.140.81.247",
			Latency:     67,
			Plugins: map[string]PeerPlugin{
				"@mainsail/webhooks": {
					Port:               4004,
					Enabled:            false,
					EstimateTotalCount: false,
				},
				"@mainsail/api-development": {
					Port:               4006,
					Enabled:            true,
					EstimateTotalCount: false,
				},
			},
			Port: 4000,
			Ports: PeerPorts{
				"@mainsail/webhooks":        -1,
				"@mainsail/api-development": 4006,
			},
			Version: "0.0.1-evm.53",
		},
	})
}

func strPtr(s string) *string {
	return &s
}

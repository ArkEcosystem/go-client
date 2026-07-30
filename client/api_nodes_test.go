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

// Get all API nodes.
func TestApiNodesService_All(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/api-nodes", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "totalCountIsEstimate": false,
			    "count": 1,
			    "first": "/api-nodes?transform=true&limit=100&page=1",
			    "last": "/api-nodes?transform=true&limit=100&page=1",
			    "next": null,
			    "pageCount": 1,
			    "previous": null,
			    "self": "/api-nodes?transform=true&limit=100&page=1",
			    "totalCount": 1
			  },
			  "data": [{
			      "url": "http://1.2.3.4:4003",
			      "version": "2.0.0",
			      "height": 1204291,
			      "latency": 10,
			      "status": "success"
			  }]
			}`)
	})

	query := &ApiNodesQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.ApiNodes.All(context.Background(), query)
	testGeneralError(t, "ApiNodes.All", err)
	testResponseUrl(t, "ApiNodes.All", response, "/api-nodes")
	testResponseStruct(t, "ApiNodes.All", responseStruct, &ApiNodesResponse{
		Meta: Meta{
			TotalCountIsEstimate: false,
			Count:                1,
			First:                "/api-nodes?transform=true&limit=100&page=1",
			Last:                 "/api-nodes?transform=true&limit=100&page=1",
			Next:                 nil,
			PageCount:            1,
			Previous:             nil,
			Self:                 "/api-nodes?transform=true&limit=100&page=1",
			TotalCount:           1,
		},
		Data: []ApiNode{{
			Url:     "http://1.2.3.4:4003",
			Version: "2.0.0",
			Height:  1204291,
			Latency: 10,
			Status:  "success",
		}},
	})
}

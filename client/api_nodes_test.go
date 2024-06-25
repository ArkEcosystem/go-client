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
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/api-nodes?page=1&limit=1",
			    "first": "/api/api-nodes?page=1&limit=1",
			    "last": "/api/api-nodes?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummy",
			      "ip": "127.0.0.1",
			      "port": 4003,
			      "version": "2.0.0"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.ApiNodes.All(context.Background(), query)
	testGeneralError(t, "ApiNodes.All", err)
	testResponseUrl(t, "ApiNodes.All", response, "/api/api-nodes")
	testResponseStruct(t, "ApiNodes.All", responseStruct, &ApiNodesResponse{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/api-nodes?page=1&limit=1",
			First:      "/api/api-nodes?page=1&limit=1",
			Last:       "/api/api-nodes?page=1&limit=1",
		},
		Data: []ApiNode{
			{
				Id:      "dummy",
				Ip:      "127.0.0.1",
				Port:    4003,
				Version: "2.0.0",
			},
		},
	})
}

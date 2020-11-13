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

// Get all entities.
func TestEntitiesService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/entities", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/entities?page=1&limit=1",
			    "first": "/api/entities?page=1&limit=1",
			    "last": "/api/entities?page=1&limit=1"
			  },
			  "data": [
			    {
			      "id": "dummyId",
			      "address": "dummy",
			      "publicKey": "dummyPublicKey",
			      "isResigned": false,
				  "type": 0,
				  "subType": 2,
			      "data": {
			        "name": "dummy",
			        "ipfsData": "dummy",
			      },
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Entities.List(context.Background(), query)
	testGeneralError(t, "Entities.List", err)
	testResponseUrl(t, "Entities.List", response, "/api/entities")
	testResponseStruct(t, "Entities.List", responseStruct, &Entities{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/entities?page=1&limit=1",
			First:      "/api/entities?page=1&limit=1",
			Last:       "/api/entities?page=1&limit=1",
		},
		Data: []Entity{{
			Id:         "dummyId",
			Address:    "dummy",
			PublicKey:  "dummyPublicKey",
			IsResigned: false,
			Type:       0,
			SubType:    2,
			Data: Data{
				Name:     "dummy",
				IpfsData: "dummy",
			},
		}},
	})
}

// Get an entity by the given id
func TestEntitiesService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/entities/dummyId", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": [
			    {
			      "id": "dummyId",
			      "address": "dummy",
			      "publicKey": "dummyPublicKey",
			      "isResigned": false,
				  "type": 0,
				  "subType": 2,
			      "data": {
			        "name": "dummy",
			        "ipfsData": "dummy",
			      },
			    }
			}`)
	})

	responseStruct, response, err := client.Entities.Get(context.Background(), "dummyId")
	testGeneralError(t, "Entities.Get", err)
	testResponseUrl(t, "Entities.Get", response, "/entities/dummyId")
	testResponseStruct(t, "Entities.Get", responseStruct, &GetLock{
		Data: Entity{
			Id:         "dummyId",
			Address:    "dummy",
			PublicKey:  "dummyPublicKey",
			IsResigned: false,
			Type:       0,
			SubType:    2,
			Data: Data{
				Name:     "dummy",
				IpfsData: "dummy",
			},
		},
	})
}

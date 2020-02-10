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

// Get all bridgechains.
func TestBridgechainsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/bridgechains", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
				"meta": {
				  "count": 1,
				  "pageCount": 1,
				  "totalCount": 1,
				  "next": null,
				  "previous": null,
				  "self": "/api/bridgechains?page=1&limit=1",
				  "first": "/api/bridgechains?page=1&limit=1",
				  "last": "/api/bridgechains?page=1&limit=1"
				},
				"data": [
				  {
				    "bridgechainId": 1,
				    "businessId": 1,
				    "name": "dummyName",
				    "seedNodes": [
				      "1.1.1.1"
				    ],
				    "genesisHash": "dummyGenesisHash",
				    "bridgechainRepository": "dummyBridgechainRepository"
				  }
				]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Bridgechains.List(context.Background(), query)
	testGeneralError(t, "Bridgechains.List", err)
	testResponseUrl(t, "Bridgechains.List", response, "/api/bridgechains")
	testResponseStruct(t, "Bridgechains.List", responseStruct, &Bridgechains{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/bridgechains?page=1&limit=1",
			First:      "/api/bridgechains?page=1&limit=1",
			Last:       "/api/bridgechains?page=1&limit=1",
		},
		Data: []Bridgechain{{
			BridgechainId: 1,
			BusinessId:    1,
			Name:          "dummyName",
			SeedNodes: []string{
				"1.1.1.1",
			},
			GenesisHash:           "dummyGenesisHash",
			BridgechainRepository: "dummyBridgechainRepository",
		}},
	})
}

// Get a bridgechain by the given id
func TestBridgechainsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/bridgechains/1", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "bridgechainId": 1,
			    "businessId": 1,
			    "name": "dummyName",
			    "seedNodes": [
			      "1.1.1.1"
			    ],
			    "genesisHash": "dummyGenesisHash",
			    "bridgechainRepository": "dummyBridgechainRepository"
			  }
			}`)
	})

	responseStruct, response, err := client.Bridgechains.Get(context.Background(), 1)
	testGeneralError(t, "Bridgechains.Get", err)
	testResponseUrl(t, "Bridgechains.Get", response, "/bridgechains/1")
	testResponseStruct(t, "Bridgechains.Get", responseStruct, &GetBridgechain{
		Data: Bridgechain{
			BridgechainId: 1,
			BusinessId:    1,
			Name:          "dummyName",
			SeedNodes: []string{
				"1.1.1.1",
			},
			GenesisHash:           "dummyGenesisHash",
			BridgechainRepository: "dummyBridgechainRepository",
		},
	})
}

// Filter all bridgechains by the given criteria.
func TestBridgechainsService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/bridgechains/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
				"meta": {
				  "count": 1,
				  "pageCount": 1,
				  "totalCount": 1,
				  "next": null,
				  "previous": null,
				  "self": "/api/bridgechains/search?page=1&limit=1",
				  "first": "/api/bridgechains/search?page=1&limit=1",
				  "last": "/api/bridgechains/search?page=1&limit=1"
				},
				"data": [
				  {
				    "bridgechainId": 1,
				    "businessId": 1,
				    "name": "dummyName",
				    "seedNodes": [
				      "1.1.1.1"
				    ],
				    "genesisHash": "dummyGenesisHash",
				    "bridgechainRepository": "dummyBridgechainRepository"
				  }
				]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &BridgechainsSearchRequest{BridgechainId: 1}
	responseStruct, response, err := client.Bridgechains.Search(context.Background(), query, body)
	testGeneralError(t, "Bridgechains.Search", err)
	testResponseUrl(t, "Bridgechains.Search", response, "/api/bridgechains/search")
	testResponseStruct(t, "Bridgechains.Search", responseStruct, &Bridgechains{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/bridgechains/search?page=1&limit=1",
			First:      "/api/bridgechains/search?page=1&limit=1",
			Last:       "/api/bridgechains/search?page=1&limit=1",
		},
		Data: []Bridgechain{{
			BridgechainId: 1,
			BusinessId:    1,
			Name:          "dummyName",
			SeedNodes: []string{
				"1.1.1.1",
			},
			GenesisHash:           "dummyGenesisHash",
			BridgechainRepository: "dummyBridgechainRepository",
		}},
	})
}

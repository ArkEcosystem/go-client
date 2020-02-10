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

// Get all businesses.
func TestBusinessesService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/businesses", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/businesses?page=1&limit=1",
			    "first": "/api/businesses?page=1&limit=1",
			    "last": "/api/businesses?page=1&limit=1"
			  },
			  "data": [
			    {
			      "businessId": 1,
			      "name": "dummyName",
			      "website": "http://dummy.website",
			      "vat": "dummyVat",
			      "repository": "dummyRepository"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Businesses.List(context.Background(), query)
	testGeneralError(t, "Businesses.List", err)
	testResponseUrl(t, "Businesses.List", response, "/api/businesses")
	testResponseStruct(t, "Businesses.List", responseStruct, &Businesses{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/businesses?page=1&limit=1",
			First:      "/api/businesses?page=1&limit=1",
			Last:       "/api/businesses?page=1&limit=1",
		},
		Data: []Business{{
			BusinessId: 1,
			Name:       "dummyName",
			Website:    "http://dummy.website",
			Vat:        "dummyVat",
			Repository: "dummyRepository",
		}},
	})
}

// Get a business by the given id
func TestBusinessesService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/businesses/1", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "businessId": 1,
			    "name": "dummyName",
			    "website": "http://dummy.website",
			    "vat": "dummyVat",
			    "repository": "dummyRepository"
			  }
			}`)
	})

	responseStruct, response, err := client.Businesses.Get(context.Background(), 1)
	testGeneralError(t, "Businesses.Get", err)
	testResponseUrl(t, "Businesses.Get", response, "/businesses/1")
	testResponseStruct(t, "Businesses.Get", responseStruct, &GetBusiness{
		Data: Business{
			BusinessId: 1,
			Name:       "dummyName",
			Website:    "http://dummy.website",
			Vat:        "dummyVat",
			Repository: "dummyRepository",
		},
	})
}

// Get all bridgechains by the given business.
func TestBusinessesService_Bridgechains(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/businesses/1/bridgechains", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/businesses/1/bridgechains?page=1&limit=1",
			    "first": "/api/businesses/1/bridgechains?page=1&limit=1",
			    "last": "/api/businesses/1/bridgechains?page=1&limit=1"
			  },
			  "data": [
			    {
			      "bridgechainId": 1,
			      "businessId": 1,
			      "name": "dummyName",
			      "seedNodes": [
			        "1.1.1.1",
			        "1.1.1.2"
			      ],
			      "genesisHash": "dummyGenesisHash",
			      "bridgechainRepository": "dummyBridgechainRepository"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Businesses.Bridgechains(context.Background(), 1, query)
	testGeneralError(t, "Businesses.Bridgechains", err)
	testResponseUrl(t, "Businesses.Bridgechains", response, "/api/businesses/1/bridgechains")
	testResponseStruct(t, "Businesses.Bridgechains", responseStruct, &GetBusinessBridgechains{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/businesses/1/bridgechains?page=1&limit=1",
			First:      "/api/businesses/1/bridgechains?page=1&limit=1",
			Last:       "/api/businesses/1/bridgechains?page=1&limit=1",
		},
		Data: []Bridgechain{{
			BridgechainId: 1,
			BusinessId:    1,
			Name:          "dummyName",
			SeedNodes: []string{
				"1.1.1.1",
				"1.1.1.2",
			},
			GenesisHash:           "dummyGenesisHash",
			BridgechainRepository: "dummyBridgechainRepository",
		}},
	})
}

// Filter all businesses by the given criteria.
func TestBusinessesService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/businesses/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/businesses/search?page=1&limit=1",
			    "first": "/api/businesses/search?page=1&limit=1",
			    "last": "/api/businesses/search?page=1&limit=1"
			  },
			  "data": [
			    {
			      "businessId": 1,
			      "name": "dummyName",
			      "website": "http://dummy.website",
			      "vat": "dummyVat",
			      "repository": "dummyRepository"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	body := &BusinessesSearchRequest{BusinessId: 1}
	responseStruct, response, err := client.Businesses.Search(context.Background(), query, body)
	testGeneralError(t, "Businesses.Search", err)
	testResponseUrl(t, "Businesses.Search", response, "/api/businesses/search")
	testResponseStruct(t, "Businesses.Search", responseStruct, &Businesses{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       "",
			Previous:   "",
			Self:       "/api/businesses/search?page=1&limit=1",
			First:      "/api/businesses/search?page=1&limit=1",
			Last:       "/api/businesses/search?page=1&limit=1",
		},
		Data: []Business{{
			BusinessId: 1,
			Name:       "dummyName",
			Website:    "http://dummy.website",
			Vat:        "dummyVat",
			Repository: "dummyRepository",
		}},
	})
}

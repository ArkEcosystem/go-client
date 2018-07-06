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

// Get the loader status.
func TestLoaderService_Status(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/loader/status", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "loaded": true,
			  "now": 3940,
			  "blocksCount": null,
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Loader.Status(context.Background())
	testGeneralError(t, "Loader.Status", err)
	testResponseUrl(t, "Loader.Status", response, "/api/loader/status")
	testResponseStruct(t, "Loader.Status", responseStruct, &LoaderStatus{
		Loaded:      true,
		Now:         3940,
		BlocksCount: 0,
		Success:     true,
	})
}

// Get the loader syncing status.
func TestLoaderService_SyncStatus(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/loader/status/sync", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "syncing": false,
			  "blocks": null,
			  "height": 3940,
			  "id": "17142334154459441029",
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Loader.SyncStatus(context.Background())
	testGeneralError(t, "Loader.SyncStatus", err)
	testResponseUrl(t, "Loader.SyncStatus", response, "/api/loader/status/sync")
	testResponseStruct(t, "Loader.SyncStatus", responseStruct, &LoaderSync{
		Syncing: false,
		Blocks:  0,
		Height:  3940,
		Id:      "17142334154459441029",
		Success: true,
	})
}

// Get the loader configuration.
func TestLoaderService_AutoConfigure(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/loader/autoconfigure", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "network": {
			    "nethash": "d9acd04bde4234a81addb8482333b4ac906bed7be5a9970ce8ada428bd083192",
			    "token": "TARK",
			    "symbol": "TѦ",
			    "explorer": "http://texplorer.ark.io",
			    "version": 23,
			    "ports": {
			      "@arkecosystem/core-p2p": 4000,
			      "@arkecosystem/core-api": 4003,
			      "@arkecosystem/core-graphql": 4005
			    },
			    "feeStatistics": [
			      {
			        "type": 0,
			        "fees": {
			          "minFee": 10000000,
			          "maxFee": 10000000,
			          "avgFee": 10000000
			        }
			      }
			    ]
			  },
			  "success": true
			}`)
	})

	responseStruct, response, err := client.Loader.AutoConfigure(context.Background())
	testGeneralError(t, "Loader.AutoConfigure", err)
	testResponseUrl(t, "Loader.AutoConfigure", response, "/api/loader/autoconfigure")
	testResponseStruct(t, "Loader.AutoConfigure", responseStruct, &LoaderAutoConfigure{
		Network: LoaderNetwork{
			Nethash:  "d9acd04bde4234a81addb8482333b4ac906bed7be5a9970ce8ada428bd083192",
			Token:    "TARK",
			Symbol:   "TѦ",
			Explorer: "http://texplorer.ark.io",
			Version:  23,
			Ports: map[string]uint16{
				"@arkecosystem/core-p2p":     4000,
				"@arkecosystem/core-api":     4003,
				"@arkecosystem/core-graphql": 4005,
			},
			FeeStatistics: []FeeStatistic{{
				Type: 0,
				Fees: Fees{
					MinFee: 10000000,
					MaxFee: 10000000,
					MvgFee: 10000000,
				},
			}},
		},
		Success: true,
	})
}

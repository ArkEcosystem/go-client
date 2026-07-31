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

// Get blockchain information.
func TestBlockchainService_Info(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/blockchain", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "block": {
			      "number": 1213625,
			      "hash": "e7fd8d878860f4f41c1f53eef8ea75b585b22016b8d3aaf2d6742c1b59894016"
			    },
			    "supply": "12727605199999969"
			  }
			}`)
	})

	responseStruct, response, err := client.Blockchain.Info(context.Background())
	testGeneralError(t, "Blockchain.Info", err)
	testResponseUrl(t, "Blockchain.Info", response, "/blockchain")
	testResponseStruct(t, "Blockchain.Info", responseStruct, &BlockchainInfo{
		Data: BlockchainData{
			Block: BlockchainBlock{
				Number: 1213625,
				Hash:   "e7fd8d878860f4f41c1f53eef8ea75b585b22016b8d3aaf2d6742c1b59894016",
			},
			Supply: newBigInt(12727605199999969),
		},
	})
}

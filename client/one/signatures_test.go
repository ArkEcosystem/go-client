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

// Get the second signature registration fee.
func TestSignaturesService_Fee(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/signatures/fee", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "fee": 10
			}`)
	})

	responseStruct, response, err := client.Signatures.Fee(context.Background())
	testGeneralError(t, "Signatures.Fee", err)
	testResponseUrl(t, "Signatures.Fee", response, "/api/signatures/fee")
	testResponseStruct(t, "Signatures.Fee", responseStruct, &SignaturesFee{
		Success: true,
		Fee:     10,
	})
}

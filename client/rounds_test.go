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

// Get the forging delegates of a round by the given id.
func TestRoundsService_Delegates(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/rounds/12345/delegates", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": [
			    {
			      "publicKey": "03ffc17c5528d490b045a9b710c754e00a536d05d9b0b78a9baa0533a246dcd98c",
			      "votes": "156947252547993"
			    }
			  ]
			}`)
	})

	responseStruct, response, err := client.Rounds.Delegates(context.Background(), 12345)
	testGeneralError(t, "Rounds.Delegates", err)
	testResponseUrl(t, "Rounds.Delegates", response, "/rounds/12345/delegates")
	testResponseStruct(t, "Rounds.Delegates", responseStruct, &GetDelegates{
		Data: []RoundDelegate{{
			PublicKey: "03ffc17c5528d490b045a9b710c754e00a536d05d9b0b78a9baa0533a246dcd98c",
			Votes:     "156947252547993",
		}},
	})
}

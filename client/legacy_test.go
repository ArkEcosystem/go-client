package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all legacy cold wallets.
func TestLegacyService_ColdWallets(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/legacy/cold-wallets", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/legacy/cold-wallets?page=1&limit=1",
			    "first": "/api/legacy/cold-wallets?page=1&limit=1",
			    "last": "/api/legacy/cold-wallets?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4",
			      "balance": "200000000000000000000",
			      "attributes": {
			        "legacyNonce": "0"
			      },
			      "mergeInfoWalletAddress": null,
			      "mergeInfoTransactionHash": null
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Legacy.ColdWallets(context.Background(), query)
	testGeneralError(t, "Legacy.ColdWallets", err)
	testResponseUrl(t, "Legacy.ColdWallets", response, "/api/legacy/cold-wallets")
	testResponseStruct(t, "Legacy.ColdWallets", responseStruct, &LegacyColdWallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/legacy/cold-wallets?page=1&limit=1",
			First:      "/api/legacy/cold-wallets?page=1&limit=1",
			Last:       "/api/legacy/cold-wallets?page=1&limit=1",
		},
		Data: []LegacyColdWallet{{
			Address: "14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4",
			Balance: newBigIntFromString("200000000000000000000"),
			Attributes: LegacyColdWalletAttributes{
				LegacyNonce: "0",
			},
			MergeInfoWalletAddress:   nil,
			MergeInfoTransactionHash: nil,
		}},
	})
}

// Get a legacy cold wallet by address.
func TestLegacyService_ColdWallet(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/legacy/cold-wallets/14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "address": "14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4",
			    "balance": "200000000000000000000",
			    "attributes": {
			      "legacyNonce": "0"
			    },
			    "mergeInfoWalletAddress": null,
			    "mergeInfoTransactionHash": null
			  }
			}`)
	})

	responseStruct, response, err := client.Legacy.ColdWallet(context.Background(), "14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4")
	testGeneralError(t, "Legacy.ColdWallet", err)
	testResponseUrl(t, "Legacy.ColdWallet", response, "/api/legacy/cold-wallets/14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4")
	testResponseStruct(t, "Legacy.ColdWallet", responseStruct, &GetLegacyColdWallet{
		Data: LegacyColdWallet{
			Address: "14wcDsexEop8Fzrnv7m4DsGgrTdQGB49Y4",
			Balance: newBigIntFromString("200000000000000000000"),
			Attributes: LegacyColdWalletAttributes{
				LegacyNonce: "0",
			},
			MergeInfoWalletAddress:   nil,
			MergeInfoTransactionHash: nil,
		},
	})
}

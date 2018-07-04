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

// Get all accounts.
func TestDelegatesService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "delegates": [
			    {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy",
			      "vote": "dummy",
			      "producedblocks": 10,
			      "missedblocks": 10,
			      "rate": 1,
			      "approval": 0.10,
			      "productivity": 0.15
			    }
			  ]
			}`)
	})

	query := &GetDelegatesQuery{Limit: 1}
	responseStruct, response, err := client.Delegates.List(context.Background(), query)
	testGeneralError(t, "Delegates.List", err)
	testResponseUrl(t, "Delegates.List", response, "/api/delegates")
	testResponseStruct(t, "Delegates.List", responseStruct, &AccountDelegates{
		Success: true,
		Delegates: []Delegate{{
			Username:       "dummy",
			Address:        "dummy",
			PublicKey:      "dummy",
			Vote:           "dummy",
			ProducedBlocks: 10,
			MissedBlocks:   10,
			Rate:           1,
			Approval:       0.10,
			Productivity:   0.15,
		}},
	})
}

// Get a delegate by the given id.
func TestDelegatesService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/get", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "delegate": {
			    "username": "dummy",
			    "address": "dummy",
			    "publicKey": "dummy",
			    "vote": "dummy",
			    "producedblocks": 10,
			    "missedblocks": 10,
			    "rate": 1,
			    "approval": 0.10,
			    "productivity": 0.15
			  }
			}`)
	})

	query := &GetDelegateQuery{Username: "dummy"}
	responseStruct, response, err := client.Delegates.Get(context.Background(), query)
	testGeneralError(t, "Delegates.Get", err)
	testResponseUrl(t, "Delegates.Get", response, "/api/delegates/get")
	testResponseStruct(t, "Delegates.Get", responseStruct, &GetDelegate{
		Success: true,
		Delegate: Delegate{
			Username:       "dummy",
			Address:        "dummy",
			PublicKey:      "dummy",
			Vote:           "dummy",
			ProducedBlocks: 10,
			MissedBlocks:   10,
			Rate:           1,
			Approval:       0.10,
			Productivity:   0.15,
		},
	})
}

// Count all delegates.
func TestDelegatesService_Count(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/count", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "count": 10
			}`)
	})

	responseStruct, response, err := client.Delegates.Count(context.Background())
	testGeneralError(t, "Delegates.Count", err)
	testResponseUrl(t, "Delegates.Count", response, "/api/delegates/count")
	testResponseStruct(t, "Delegates.Count", responseStruct, &DelegatesCount{
		Success: true,
		Count:   10,
	})
}

// Get the delegate registration fee.
func TestDelegatesService_Fee(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/fee", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "fee": 10
			}`)
	})

	responseStruct, response, err := client.Delegates.Fee(context.Background())
	testGeneralError(t, "Delegates.Fee", err)
	testResponseUrl(t, "Delegates.Fee", response, "/api/delegates/fee")
	testResponseStruct(t, "Delegates.Fee", responseStruct, &DelegateFee{
		Success: true,
		Fee:     10,
	})
}

// Get the forged totals by the given public key.
func TestDelegatesService_ForgedByAccount(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/forging/getForgedByAccount", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "fees": "10",
			  "rewards": "10",
			  "forged": "10"
			}`)
	})

	responseStruct, response, err := client.Delegates.ForgedByAccount(context.Background(), "dummy")
	testGeneralError(t, "Delegates.ForgedByAccount", err)
	testResponseUrl(t, "Delegates.ForgedByAccount", response, "/api/delegates/forging/getForgedByAccount")
	testResponseStruct(t, "Delegates.ForgedByAccount", responseStruct, &ForgedByDelegate{
		Success: true,
		Fees:    "10",
		Rewards: "10",
		Forged:  "10",
	})
}

// Filter all delegates by the given criteria.
func TestDelegatesService_Search(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/search", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "delegates": [
			    {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy",
			      "vote": "dummy",
			      "producedblocks": 10,
			      "missedblocks": 10,
			      "rate": 1,
			      "approval": 0.10,
			      "productivity": 0.15
			    }
			  ]
			}`)
	})

	query := &DelegateSearchQuery{Limit: 1, Q: "dummy"}
	responseStruct, response, err := client.Delegates.Search(context.Background(), query)
	testGeneralError(t, "Delegates.Search", err)
	testResponseUrl(t, "Delegates.Search", response, "/api/delegates/search")
	testResponseStruct(t, "Delegates.Search", responseStruct, &AccountDelegates{
		Success: true,
		Delegates: []Delegate{{
			Username:       "dummy",
			Address:        "dummy",
			PublicKey:      "dummy",
			Vote:           "dummy",
			ProducedBlocks: 10,
			MissedBlocks:   10,
			Rate:           1,
			Approval:       0.10,
			Productivity:   0.15,
		}},
	})
}

// Get all voters by the given public key.
func TestDelegatesService_Voters(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/voters", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "accounts": [
			    {
			      "username": "dummy",
			      "address": "dummy",
			      "publicKey": "dummy",
			      "balance": "dummy"
			    }
			  ]
			}`)
	})

	responseStruct, response, err := client.Delegates.Voters(context.Background(), "dummy")
	testGeneralError(t, "Delegates.Voters", err)
	testResponseUrl(t, "Delegates.Voters", response, "/api/delegates/voters")
	testResponseStruct(t, "Delegates.Voters", responseStruct, &Accounts{
		Success: true,
		Accounts: []Account{{
			Username:  "dummy",
			Address:   "dummy",
			PublicKey: "dummy",
			Balance:   "dummy",
		}},
	})
}

// Get a list of the next forgers.
func TestDelegatesService_NextForgers(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/delegates/getNextForgers", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "success": true,
			  "currentBlock": 10,
			  "currentSlot": 10,
			  "delegates": [
			    "dummy_1",
			    "dummy_2",
			    "dummy_3",
			    "dummy_4",
			    "dummy_5",
			    "dummy_6",
			    "dummy_7",
			    "dummy_8",
			    "dummy_9",
			    "dummy_10"
			  ]
			}`)
	})

	responseStruct, response, err := client.Delegates.NextForgers(context.Background())
	testGeneralError(t, "Delegates.NextForgers", err)
	testResponseUrl(t, "Delegates.NextForgers", response, "/api/delegates/getNextForgers")
	testResponseStruct(t, "Delegates.NextForgers", responseStruct, &NextForger{
		Success:      true,
		CurrentBlock: 10,
		CurrentSlot:  10,
		Delegates: []string{
			"dummy_1",
			"dummy_2",
			"dummy_3",
			"dummy_4",
			"dummy_5",
			"dummy_6",
			"dummy_7",
			"dummy_8",
			"dummy_9",
			"dummy_10",
		},
	})
}

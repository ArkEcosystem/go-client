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



// TestRoundsService_All tests the All method.
func TestRoundsService_All(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/rounds", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{
			"meta": {
				"totalCountIsEstimate": false,
				"count": 100,
				"first": "/rounds?limit=100&page=1",
				"last": "/rounds?limit=100&page=230",
								"next": "/rounds?limit=100&page=2",
				"pageCount": 230,
				"next": null,
				"previous": null,
				"self": "/rounds?limit=100&page=1",
				"totalCount": 22908
			},
			"data": [
				{
					"round": "22908",
					"roundHeight": "1214072",
					"validators": [
						"0270d7400c6262ea35dadbb2f71d6e871dc48c3f1869fb68e03f9a0179b530226c",
						"03d878ed802a26f299c1492e5a5db7615a45bfa05cc703881d7f28286ce8bb96a4"
					],
					"votes": [
						"240245486603773",
						"240240776603773"
					]
				}
			]
		}`)
	})

	query := &Pagination{Limit: 100}
	responseStruct, response, err := client.Rounds.All(context.Background(), query)
	testGeneralError(t, "Rounds.All", err)
	testResponseUrl(t, "Rounds.All", response, "/rounds?limit=100&page=1")
	testResponseStruct(t, "Rounds.All", responseStruct, &GetRounds{
		Meta: Meta{
			TotalCountIsEstimate: false,
			Count:                100,
			First:                "/rounds?limit=100&page=1",
			Last:                 "/rounds?limit=100&page=230",
			Next:       nil,
			PageCount:            230,
			Previous:             nil,
			Self:                 "/rounds?limit=100&page=1",
			TotalCount:           22908,
		},
		Data: []RoundData{{
			Round:       "22908",
			RoundHeight: "1214072",
			Validators: []string{
				"0270d7400c6262ea35dadbb2f71d6e871dc48c3f1869fb68e03f9a0179b530226c",
				"03d878ed802a26f299c1492e5a5db7615a45bfa05cc703881d7f28286ce8bb96a4",
			},
			Votes: []string{
				"240245486603773",
				"240240776603773",
			},
		}},
	})
}

// TestRoundsService_Show tests the Show method.
func TestRoundsService_Show(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/rounds/12345", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{
			"data": {
				"round": "22908",
				"roundHeight": "1214072",
				"validators": [
					"0270d7400c6262ea35dadbb2f71d6e871dc48c3f1869fb68e03f9a0179b530226c",
					"03d878ed802a26f299c1492e5a5db7615a45bfa05cc703881d7f28286ce8bb96a4"
				],
				"votes": [
					"240245486603773",
					"240240776603773"
				]
			}
		}`)
	})

	responseStruct, response, err := client.Rounds.Show(context.Background(), 12345)
	testGeneralError(t, "Rounds.Show", err)
	testResponseUrl(t, "Rounds.Show", response, "/rounds/12345")
	testResponseStruct(t, "Rounds.Show", responseStruct, &GetRound{
		Data: RoundData{
			Round:       "22908",
			RoundHeight: "1214072",
			Validators: []string{
				"0270d7400c6262ea35dadbb2f71d6e871dc48c3f1869fb68e03f9a0179b530226c",
				"03d878ed802a26f299c1492e5a5db7615a45bfa05cc703881d7f28286ce8bb96a4",
			},
			Votes: []string{
				"240245486603773",
				"240240776603773",
			},
		},
	})
}


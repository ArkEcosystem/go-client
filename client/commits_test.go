package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// TestCommitsService_GetCommit tests the GetCommit method.
func TestCommitsService_GetCommit(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/commits/1", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer, `{
			"data": {
				"blockNumber": "1",
				"signature": "819b0dc4bcb27a95e9c0c83fe5294cdad8e4828dda7dfb050dc1f723ad5fef283855db194a64e2754222086c76291b04115ac64702c47c9794c5171b0c7635b68651021763818400259ed9f32fcc9051efadf751efcf70ca45de88cd7811e5fb",
				"validators": [
					"02ba352eab5ddfd1a19f5681162b1550b035003ed01bd83cd173f48dbb8ebb2907",
					"02ab5bebf52333bdad0d3d75e6b5b1f602adf7b8af6faa87c23b2e121318f95542",
					"03f3f6d09101d6fd97cc5b987707442c14cb4b990b097852f265ef41631ae4c7d5"
				]
			}
		}`)
	})

	responseStruct, response, err := client.Commits.GetCommit(context.Background(), 1)

	testGeneralError(t, "Commits.GetCommit", err)
	testResponseUrl(t, "Commits.GetCommit", response, "/commits/1")
	testResponseStruct(t, "Commits.GetCommit", responseStruct, &CommitResponse{
		Data: CommitData{
			BlockNumber: "1",
			Signature:   "819b0dc4bcb27a95e9c0c83fe5294cdad8e4828dda7dfb050dc1f723ad5fef283855db194a64e2754222086c76291b04115ac64702c47c9794c5171b0c7635b68651021763818400259ed9f32fcc9051efadf751efcf70ca45de88cd7811e5fb",
			Validators: []string{
				"02ba352eab5ddfd1a19f5681162b1550b035003ed01bd83cd173f48dbb8ebb2907",
				"02ab5bebf52333bdad0d3d75e6b5b1f602adf7b8af6faa87c23b2e121318f95542",
				"03f3f6d09101d6fd97cc5b987707442c14cb4b990b097852f265ef41631ae4c7d5",
			},
		},
	})
}

package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestContractsService_All(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/contracts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "consensus": {
			      "activeImplementation": "0x522B3294E6d06aA25Ad0f1B8891242E335D3B459",
			      "address": "0x535B3D7A252fa034Ed71F0C53ec0C6F784cB64E1",
			      "implementations": ["0x522B3294E6d06aA25Ad0f1B8891242E335D3B459"],
			      "proxy": "UUPS"
			    }
			  }
			}`)
	})

	responseStruct, response, err := client.Contracts.All(context.Background())
	testGeneralError(t, "Contracts.All", err)
	testResponseUrl(t, "Contracts.All", response, "/api/contracts")
	testResponseStruct(t, "Contracts.All", responseStruct, &ContractsResponse{
		Data: map[string]Contract{
			"consensus": {
				ActiveImplementation: "0x522B3294E6d06aA25Ad0f1B8891242E335D3B459",
				Address:              "0x535B3D7A252fa034Ed71F0C53ec0C6F784cB64E1",
				Implementations:      []string{"0x522B3294E6d06aA25Ad0f1B8891242E335D3B459"},
				Proxy:                "UUPS",
			},
		},
	})
}

func TestContractsService_Abi(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/contracts/consensus/0x522B3294E6d06aA25Ad0f1B8891242E335D3B459/abi", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "abi": [
			      {
			        "name": "activeValidatorsCount",
			        "type": "function",
			        "inputs": [],
			        "outputs": [{"name": "", "type": "uint256", "internalType": "uint256"}],
			        "stateMutability": "view"
			      }
			    ]
			  }
			}`)
	})

	responseStruct, response, err := client.Contracts.Abi(context.Background(), "consensus", "0x522B3294E6d06aA25Ad0f1B8891242E335D3B459")
	testGeneralError(t, "Contracts.Abi", err)
	testResponseUrl(t, "Contracts.Abi", response, "/api/contracts/consensus/0x522B3294E6d06aA25Ad0f1B8891242E335D3B459/abi")
	testResponseStruct(t, "Contracts.Abi", responseStruct, &ContractAbiResponse{
		Data: ContractAbi{
			Abi: []interface{}{
				map[string]interface{}{
					"name":   "activeValidatorsCount",
					"type":   "function",
					"inputs": []interface{}{},
					"outputs": []interface{}{
						map[string]interface{}{"name": "", "type": "uint256", "internalType": "uint256"},
					},
					"stateMutability": "view",
				},
			},
		},
	})
}

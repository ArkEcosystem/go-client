package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// Get all tokens.
func TestTokensService_All(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/tokens?page=1&limit=1",
			    "first": "/api/tokens?page=1&limit=1",
			    "last": "/api/tokens?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "0x12f6677522292654a231007c47b07971a7610908",
			      "symbol": "DARK20",
			      "name": "DARK20",
			      "decimals": 18,
			      "totalSupply": "100000000000000000000000000",
			      "deploymentHash": "cba853a2d1fcbad4c9a22f88c9a5f681425b18a6b6df4981a7e2f83176c4c1ce"
			    }
			  ]
			}`)
	})

	query := &TokensQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Tokens.All(context.Background(), query)
	testGeneralError(t, "Tokens.All", err)
	testResponseUrl(t, "Tokens.All", response, "/api/tokens")
	testResponseStruct(t, "Tokens.All", responseStruct, &Tokens{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/tokens?page=1&limit=1",
			First:      "/api/tokens?page=1&limit=1",
			Last:       "/api/tokens?page=1&limit=1",
		},
		Data: []Token{{
			Address:        "0x12f6677522292654a231007c47b07971a7610908",
			Symbol:         "DARK20",
			Name:           "DARK20",
			Decimals:       18,
			TotalSupply:    newBigIntFromString("100000000000000000000000000"),
			DeploymentHash: "cba853a2d1fcbad4c9a22f88c9a5f681425b18a6b6df4981a7e2f83176c4c1ce",
		}},
	})
}

func TestTokensService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/0x12f6677522292654a231007c47b07971a7610908", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "address": "0x12f6677522292654a231007c47b07971a7610908",
			    "symbol": "DARK20",
			    "name": "DARK20",
			    "decimals": 18,
			    "totalSupply": "100000000000000000000000000",
			    "deploymentHash": "cba853a2d1fcbad4c9a22f88c9a5f681425b18a6b6df4981a7e2f83176c4c1ce"
			  }
			}`)
	})

	responseStruct, response, err := client.Tokens.Get(context.Background(), "0x12f6677522292654a231007c47b07971a7610908")
	testGeneralError(t, "Tokens.Get", err)
	testResponseUrl(t, "Tokens.Get", response, "/api/tokens/0x12f6677522292654a231007c47b07971a7610908")
	testResponseStruct(t, "Tokens.Get", responseStruct, &GetToken{
		Data: Token{
			Address:        "0x12f6677522292654a231007c47b07971a7610908",
			Symbol:         "DARK20",
			Name:           "DARK20",
			Decimals:       18,
			TotalSupply:    newBigIntFromString("100000000000000000000000000"),
			DeploymentHash: "cba853a2d1fcbad4c9a22f88c9a5f681425b18a6b6df4981a7e2f83176c4c1ce",
		},
	})
}

func TestTokensService_Transfers(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/transfers", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/tokens/transfers?page=1&limit=1",
			    "first": "/api/tokens/transfers?page=1&limit=1",
			    "last": "/api/tokens/transfers?page=1&limit=1"
			  },
			  "data": [
			    {
			      "transactionHash": "371fbdd1d9e5f0489d8bc102f25eff50ddb46e01b56a9cc0d205e3aa153ba01e",
			      "from": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "to": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "value": "123000000000000000000",
			      "functionSig": "0x4885b254",
			      "blockNumber": "23072686",
			      "timestamp": "1784046832110",
			      "token": {
			        "address": "0x180a864a755fed0144c622df49b83db577befefb",
			        "name": "DARK20",
			        "symbol": "DARK20",
			        "decimals": 18
			      }
			    }
			  ]
			}`)
	})

	query := &TokenTransfersQuery{TokenLookupQuery: TokenLookupQuery{Pagination: Pagination{Limit: 1}}}
	responseStruct, response, err := client.Tokens.Transfers(context.Background(), query)
	testGeneralError(t, "Tokens.Transfers", err)
	testResponseUrl(t, "Tokens.Transfers", response, "/api/tokens/transfers")
	testResponseStruct(t, "Tokens.Transfers", responseStruct, &TokenActions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/tokens/transfers?page=1&limit=1",
			First:      "/api/tokens/transfers?page=1&limit=1",
			Last:       "/api/tokens/transfers?page=1&limit=1",
		},
		Data: []TokenAction{{
			TransactionHash: "371fbdd1d9e5f0489d8bc102f25eff50ddb46e01b56a9cc0d205e3aa153ba01e",
			From:            "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			To:              "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			Value:           newBigIntFromString("123000000000000000000"),
			FunctionSig:     "0x4885b254",
			BlockNumber:     "23072686",
			Timestamp:       "1784046832110",
			Token: TokenActionToken{
				Address:  "0x180a864a755fed0144c622df49b83db577befefb",
				Name:     "DARK20",
				Symbol:   "DARK20",
				Decimals: 18,
			},
		}},
	})
}

func TestTokensService_Approvals(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/approvals", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/tokens/approvals?page=1&limit=1",
			    "first": "/api/tokens/approvals?page=1&limit=1",
			    "last": "/api/tokens/approvals?page=1&limit=1"
			  },
			  "data": [
			    {
			      "transactionHash": "88e074a73ebc73322e296039ba890682715d91f53b79034b45adb958d4c6081e",
			      "from": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "to": "0x5a223F4434D5Bd8478100EEb3b0166a57A26350d",
			      "value": "246000000000000000000",
			      "functionSig": "0x095ea7b3",
			      "blockNumber": "23072683",
			      "timestamp": "1784046808110",
			      "token": {
			        "address": "0x180a864a755fed0144c622df49b83db577befefb",
			        "name": "DARK20",
			        "symbol": "DARK20",
			        "decimals": 18
			      }
			    }
			  ]
			}`)
	})

	query := &TokenApprovalsQuery{TokenLookupQuery: TokenLookupQuery{Pagination: Pagination{Limit: 1}}}
	responseStruct, response, err := client.Tokens.Approvals(context.Background(), query)
	testGeneralError(t, "Tokens.Approvals", err)
	testResponseUrl(t, "Tokens.Approvals", response, "/api/tokens/approvals")
	testResponseStruct(t, "Tokens.Approvals", responseStruct, &TokenActions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/tokens/approvals?page=1&limit=1",
			First:      "/api/tokens/approvals?page=1&limit=1",
			Last:       "/api/tokens/approvals?page=1&limit=1",
		},
		Data: []TokenAction{{
			TransactionHash: "88e074a73ebc73322e296039ba890682715d91f53b79034b45adb958d4c6081e",
			From:            "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			To:              "0x5a223F4434D5Bd8478100EEb3b0166a57A26350d",
			Value:           newBigIntFromString("246000000000000000000"),
			FunctionSig:     "0x095ea7b3",
			BlockNumber:     "23072683",
			Timestamp:       "1784046808110",
			Token: TokenActionToken{
				Address:  "0x180a864a755fed0144c622df49b83db577befefb",
				Name:     "DARK20",
				Symbol:   "DARK20",
				Decimals: 18,
			},
		}},
	})
}

func TestTokensService_Whitelist(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/whitelist", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/tokens/whitelist?page=1&limit=1",
			    "first": "/api/tokens/whitelist?page=1&limit=1",
			    "last": "/api/tokens/whitelist?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "0x12f6677522292654a231007c47b07971a7610908",
			      "comment": "SamCoin",
			      "createdAt": "2026-03-02T15:00:00.000Z"
			    }
			  ]
			}`)
	})

	query := &Pagination{Limit: 1}
	responseStruct, response, err := client.Tokens.Whitelist(context.Background(), query)
	testGeneralError(t, "Tokens.Whitelist", err)
	testResponseUrl(t, "Tokens.Whitelist", response, "/api/tokens/whitelist")
	testResponseStruct(t, "Tokens.Whitelist", responseStruct, &TokenWhitelist{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/tokens/whitelist?page=1&limit=1",
			First:      "/api/tokens/whitelist?page=1&limit=1",
			Last:       "/api/tokens/whitelist?page=1&limit=1",
		},
		Data: []TokenWhitelistEntry{{
			Address:   "0x12f6677522292654a231007c47b07971a7610908",
			Comment:   "SamCoin",
			CreatedAt: "2026-03-02T15:00:00.000Z",
		}},
	})
}

func TestTokensService_TransfersFor(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/0x180a864a755fed0144c622df49b83db577befefb/transfers", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "totalCountIsEstimate": false
			  },
			  "results": [
			    {
			      "transactionHash": "371fbdd1d9e5f0489d8bc102f25eff50ddb46e01b56a9cc0d205e3aa153ba01e",
			      "from": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "to": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "value": "123000000000000000000",
			      "functionSig": "0x4885b254",
			      "blockNumber": "23072686",
			      "timestamp": "1784046832110",
			      "token": {
			        "address": "0x180a864a755fed0144c622df49b83db577befefb",
			        "name": "DARK20",
			        "symbol": "DARK20",
			        "decimals": 18
			      }
			    }
			  ],
			  "totalCount": 73
			}`)
	})

	query := &TokenLookupQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Tokens.TransfersFor(context.Background(), "0x180a864a755fed0144c622df49b83db577befefb", query)
	testGeneralError(t, "Tokens.TransfersFor", err)
	testResponseUrl(t, "Tokens.TransfersFor", response, "/api/tokens/0x180a864a755fed0144c622df49b83db577befefb/transfers")
	testResponseStruct(t, "Tokens.TransfersFor", responseStruct, &TokenActionsResults{
		Meta: Meta{
			TotalCountIsEstimate: false,
		},
		Results: []TokenAction{{
			TransactionHash: "371fbdd1d9e5f0489d8bc102f25eff50ddb46e01b56a9cc0d205e3aa153ba01e",
			From:            "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			To:              "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			Value:           newBigIntFromString("123000000000000000000"),
			FunctionSig:     "0x4885b254",
			BlockNumber:     "23072686",
			Timestamp:       "1784046832110",
			Token: TokenActionToken{
				Address:  "0x180a864a755fed0144c622df49b83db577befefb",
				Name:     "DARK20",
				Symbol:   "DARK20",
				Decimals: 18,
			},
		}},
		TotalCount: 73,
	})
}

func TestTokensService_ApprovalsFor(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/0x180a864a755fed0144c622df49b83db577befefb/approvals", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "totalCountIsEstimate": false
			  },
			  "results": [
			    {
			      "transactionHash": "88e074a73ebc73322e296039ba890682715d91f53b79034b45adb958d4c6081e",
			      "from": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "to": "0x5a223F4434D5Bd8478100EEb3b0166a57A26350d",
			      "value": "246000000000000000000",
			      "functionSig": "0x095ea7b3",
			      "blockNumber": "23072683",
			      "timestamp": "1784046808110",
			      "token": {
			        "address": "0x180a864a755fed0144c622df49b83db577befefb",
			        "name": "DARK20",
			        "symbol": "DARK20",
			        "decimals": 18
			      }
			    }
			  ],
			  "totalCount": 47
			}`)
	})

	query := &TokenLookupQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Tokens.ApprovalsFor(context.Background(), "0x180a864a755fed0144c622df49b83db577befefb", query)
	testGeneralError(t, "Tokens.ApprovalsFor", err)
	testResponseUrl(t, "Tokens.ApprovalsFor", response, "/api/tokens/0x180a864a755fed0144c622df49b83db577befefb/approvals")
	testResponseStruct(t, "Tokens.ApprovalsFor", responseStruct, &TokenActionsResults{
		Meta: Meta{
			TotalCountIsEstimate: false,
		},
		Results: []TokenAction{{
			TransactionHash: "88e074a73ebc73322e296039ba890682715d91f53b79034b45adb958d4c6081e",
			From:            "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			To:              "0x5a223F4434D5Bd8478100EEb3b0166a57A26350d",
			Value:           newBigIntFromString("246000000000000000000"),
			FunctionSig:     "0x095ea7b3",
			BlockNumber:     "23072683",
			Timestamp:       "1784046808110",
			Token: TokenActionToken{
				Address:  "0x180a864a755fed0144c622df49b83db577befefb",
				Name:     "DARK20",
				Symbol:   "DARK20",
				Decimals: 18,
			},
		}},
		TotalCount: 47,
	})
}

func TestTokensService_HoldersFor(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/tokens/0x12f6677522292654a231007c47b07971a7610908/holders", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "totalCountIsEstimate": false
			  },
			  "results": [
			    {
			      "tokenAddress": "0x12f6677522292654a231007c47b07971a7610908",
			      "address": "0xA5cc0BfEB09742C5e4C610f2EBaaB82Eb142Ca10",
			      "balance": "99944601000000000000000000"
			    }
			  ],
			  "totalCount": 5
			}`)
	})

	responseStruct, response, err := client.Tokens.HoldersFor(context.Background(), "0x12f6677522292654a231007c47b07971a7610908")
	testGeneralError(t, "Tokens.HoldersFor", err)
	testResponseUrl(t, "Tokens.HoldersFor", response, "/api/tokens/0x12f6677522292654a231007c47b07971a7610908/holders")
	testResponseStruct(t, "Tokens.HoldersFor", responseStruct, &TokenHolders{
		Meta: Meta{
			TotalCountIsEstimate: false,
		},
		Results: []TokenHolder{{
			TokenAddress: "0x12f6677522292654a231007c47b07971a7610908",
			Address:      "0xA5cc0BfEB09742C5e4C610f2EBaaB82Eb142Ca10",
			Balance:      newBigIntFromString("99944601000000000000000000"),
		}},
		TotalCount: 5,
	})
}

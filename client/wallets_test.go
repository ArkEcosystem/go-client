package client

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

// Get all wallets.
func TestWalletsService_List(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets?page=1&limit=1",
			    "first": "/api/wallets?page=1&limit=1",
			    "last": "/api/wallets?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "nonce": "1",
			      "balance": "1000000000",
			      "attributes": {}
			    }
			  ]
			}`)
	})

	query := &WalletsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.List(context.Background(), query)
	testGeneralError(t, "Wallets.List", err)
	testResponseUrl(t, "Wallets.List", response, "/api/wallets")
	testResponseStruct(t, "Wallets.List", responseStruct, &Wallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets?page=1&limit=1",
			First:      "/api/wallets?page=1&limit=1",
			Last:       "/api/wallets?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:   "dummy",
			PublicKey: "dummy",
			Nonce:     newBigInt(1),
			Balance:   newBigInt(1000000000),
		}},
	})
}

// Get all wallets sorted by balance in descending order.
func TestWalletsService_Top(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/top", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/top?page=1&limit=1",
			    "first": "/api/wallets/top?page=1&limit=1",
			    "last": "/api/wallets/top?page=1&limit=1"
			  },
			  "data": [
			    {
			      "address": "dummy",
			      "publicKey": "dummy",
			      "nonce": "1",
			      "balance": "1000000000"
			    }
			  ]
			}`)
	})

	query := &WalletsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.Top(context.Background(), query)
	testGeneralError(t, "Wallets.Top", err)
	testResponseUrl(t, "Wallets.Top", response, "/api/wallets/top")
	testResponseStruct(t, "Wallets.Top", responseStruct, &Wallets{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/top?page=1&limit=1",
			First:      "/api/wallets/top?page=1&limit=1",
			Last:       "/api/wallets/top?page=1&limit=1",
		},
		Data: []Wallet{{
			Address:   "dummy",
			PublicKey: "dummy",
			Nonce:     newBigInt(1),
			Balance:   newBigInt(1000000000),
		}},
	})
}

// Get a wallet by the given id. (address, publicKey and username are valid)
func TestWalletsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "address": "dummy",
			    "publicKey": "dummy",
		      "nonce": "1",
			    "balance": "1000000000",
			    "updated_at": "2026-03-02T15:00:00.000Z",
			    "tokenCount": 2
			  }
			}`)
	})

	responseStruct, response, err := client.Wallets.Get(context.Background(), "dummy")
	testGeneralError(t, "Wallets.Get", err)
	testResponseUrl(t, "Wallets.Get", response, "/api/wallets/dummy")
	testResponseStruct(t, "Wallets.Get", responseStruct, &GetWallet{
		Data: Wallet{
			Address:    "dummy",
			PublicKey:  "dummy",
			Nonce:      newBigInt(1),
			Balance:    newBigInt(1000000000),
			UpdatedAt:  "2026-03-02T15:00:00.000Z",
			TokenCount: 2,
		},
	})
}

// Get a wallet with a Mainsail wei-scale balance that overflows uint64.
func TestWalletsService_Get_WeiScaleBalance(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "address": "dummy",
			    "publicKey": "dummy",
			    "nonce": "1",
			    "balance": "260000000000000000000"
			  }
			}`)
	})

	responseStruct, response, err := client.Wallets.Get(context.Background(), "dummy")
	testGeneralError(t, "Wallets.Get", err)
	testResponseUrl(t, "Wallets.Get", response, "/api/wallets/dummy")

	want, _ := new(big.Int).SetString("260000000000000000000", 10)
	if responseStruct.Data.Balance.Int == nil || responseStruct.Data.Balance.Cmp(want) != 0 {
		t.Errorf("[Wallets.Get][Balance] got %v, want %v", responseStruct.Data.Balance.Int, want)
	}
}

// Get all transactions for the given wallet.
func TestWalletsService_Transactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "10000000",
			      "gas": "21000",
			      "gasPrice": "10000000",
			      "senderPublicKey": "dummy",
			      "to": "dummy",
			      "from": "dummy",
			      "data": "0x",
			      "signature": "dummy",
			      "confirmations": 10,
			      "timestamp": "1719434741918",
			      "nonce": "1",
			      "receipt": {
			        "cumulativeGasUsed": 21000,
			        "gasRefunded": 0,
			        "gasUsed": 21000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &TransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.Transactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/dummy/transactions?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		}},
	})
}

// Get all transactions sent by the given wallet.
func TestWalletsService_SentTransactions(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions/sent", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions/sent?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "10000000",
			      "gas": "21000",
			      "gasPrice": "10000000",
			      "senderPublicKey": "dummy",
			      "to": "dummy",
			      "from": "dummy",
			      "data": "0x",
			      "signature": "dummy",
			      "confirmations": 10,
			      "timestamp": "1719434741918",
			      "nonce": "1",
			      "receipt": {
			        "cumulativeGasUsed": 21000,
			        "gasRefunded": 0,
			        "gasUsed": 21000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &TransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.SentTransactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions/sent")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions/sent?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions/sent?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		}},
	})
}

// Get all transactions received by the given wallet.
func TestWalletsService_ReceivedTransaction(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/transactions/received", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/transactions/received?page=1&limit=1",
			    "first": "/api/wallets/dummy/transactions/received?page=1&limit=1",
			    "last": "/api/wallets/dummy/transactions/received?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "10000000",
			      "gas": "21000",
			      "gasPrice": "10000000",
			      "senderPublicKey": "dummy",
			      "to": "dummy",
			      "from": "dummy",
			      "data": "0x",
			      "signature": "dummy",
			      "confirmations": 10,
			      "timestamp": "1719434741918",
			      "nonce": "1",
			      "receipt": {
			        "cumulativeGasUsed": 21000,
			        "gasRefunded": 0,
			        "gasUsed": 21000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &TransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.ReceivedTransactions(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Transactions", err)
	testResponseUrl(t, "Wallets.Transactions", response, "/api/wallets/dummy/transactions/received")
	testResponseStruct(t, "Wallets.Transactions", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/dummy/transactions/received?page=1&limit=1",
			First:      "/api/wallets/dummy/transactions/received?page=1&limit=1",
			Last:       "/api/wallets/dummy/transactions/received?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(10000000),
			Gas:             newBigInt(21000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 21000,
				GasRefunded:       0,
				GasUsed:           21000,
				Status:            1,
			},
		}},
	})
}

// Get all votes by the given wallet.
func TestWalletsService_Votes(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/dummy/votes", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/dummy/votes?page=1&limit=1",
			    "first": "/api/wallets/dummy/votes?page=1&limit=1",
			    "last": "/api/wallets/dummy/votes?page=1&limit=1"
			  },
			  "data": [
			    {
			      "hash": "dummy",
			      "blockHash": "dummy",
			      "value": "0",
			      "gas": "100000",
			      "gasPrice": "10000000",
			      "senderPublicKey": "dummy",
			      "to": "dummy",
			      "from": "dummy",
			      "data": "0x",
			      "signature": "dummy",
			      "confirmations": 10,
			      "timestamp": "1719434741918",
			      "nonce": "1",
			      "receipt": {
			        "cumulativeGasUsed": 100000,
			        "gasRefunded": 0,
			        "gasUsed": 100000,
			        "status": 1
			      }
			    }
			  ]
			}`)
	})

	query := &TransactionsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.Votes(context.Background(), "dummy", query)
	testGeneralError(t, "Wallets.Votes", err)
	testResponseUrl(t, "Wallets.Votes", response, "/api/wallets/dummy/votes")
	testResponseStruct(t, "Wallets.Votes", responseStruct, &Transactions{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/dummy/votes?page=1&limit=1",
			First:      "/api/wallets/dummy/votes?page=1&limit=1",
			Last:       "/api/wallets/dummy/votes?page=1&limit=1",
		},
		Data: []Transaction{{
			Hash:            "dummy",
			BlockHash:       "dummy",
			Value:           newBigInt(0),
			Gas:             newBigInt(100000),
			GasPrice:        newBigInt(10000000),
			SenderPublicKey: "dummy",
			To:              "dummy",
			From:            "dummy",
			Data:            "0x",
			Signature:       "dummy",
			Confirmations:   10,
			Timestamp:       "1719434741918",
			Nonce:           newBigInt(1),
			Receipt: TransactionReceipt{
				CumulativeGasUsed: 100000,
				GasRefunded:       0,
				GasUsed:           100000,
				Status:            1,
			},
		}},
	})
}

func TestWalletsService_TokensFor(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1",
			    "first": "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1",
			    "last": "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1"
			  },
			  "data": [
			    {
			      "tokenAddress": "0x180a864a755fed0144c622df49b83db577befefb",
			      "address": "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			      "balance": "500000000000000000000",
			      "name": "DARK20",
			      "symbol": "DARK20",
			      "decimals": 18,
			      "supply": "100000000000000000000000000"
			    }
			  ]
			}`)
	})

	query := &WalletTokensForQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Wallets.TokensFor(context.Background(), "0x29C73Db411118fa1Bf029390B097E0b23f64A496", query)
	testGeneralError(t, "Wallets.TokensFor", err)
	testResponseUrl(t, "Wallets.TokensFor", response, "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens")
	testResponseStruct(t, "Wallets.TokensFor", responseStruct, &WalletTokens{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1",
			First:      "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1",
			Last:       "/api/wallets/0x29C73Db411118fa1Bf029390B097E0b23f64A496/tokens?page=1&limit=1",
		},
		Data: []WalletToken{{
			TokenAddress: "0x180a864a755fed0144c622df49b83db577befefb",
			Address:      "0x29C73Db411118fa1Bf029390B097E0b23f64A496",
			Balance:      newBigIntFromString("500000000000000000000"),
			Name:         "DARK20",
			Symbol:       "DARK20",
			Decimals:     18,
			Supply:       newBigIntFromString("100000000000000000000000000"),
		}},
	})
}

func TestWalletsService_Tokens(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/wallets/tokens", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1",
			    "first": "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1",
			    "last": "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1"
			  },
			  "data": [
			    {
			      "token": "0x180a864a755fed0144c622df49b83db577befefb",
			      "symbol": "DARK20",
			      "name": "DARK20",
			      "decimals": 18,
			      "supply": "100000000000000000000000000",
			      "addresses": {
			        "0x29C73Db411118fa1Bf029390B097E0b23f64A496": "500000000000000000000"
			      }
			    }
			  ]
			}`)
	})

	query := &WalletTokensQuery{
		Pagination: Pagination{Limit: 1},
		Addresses:  CommaSeparated{"0x29C73Db411118fa1Bf029390B097E0b23f64A496"},
	}
	responseStruct, response, err := client.Wallets.Tokens(context.Background(), query)
	testGeneralError(t, "Wallets.Tokens", err)
	testResponseUrl(t, "Wallets.Tokens", response, "/api/wallets/tokens")
	if !strings.Contains(response.Request.URL.String(), "addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496") {
		t.Errorf("[Wallets.Tokens][URL] got %+v, want addresses param present", response.Request.URL.String())
	}
	testResponseStruct(t, "Wallets.Tokens", responseStruct, &WalletTokenAddressesResponse{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1",
			First:      "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1",
			Last:       "/api/wallets/tokens?addresses=0x29C73Db411118fa1Bf029390B097E0b23f64A496&page=1&limit=1",
		},
		Data: []WalletTokenAddresses{{
			Token:    "0x180a864a755fed0144c622df49b83db577befefb",
			Symbol:   "DARK20",
			Name:     "DARK20",
			Decimals: 18,
			Supply:   newBigIntFromString("100000000000000000000000000"),
			Addresses: map[string]BigInt{
				"0x29C73Db411118fa1Bf029390B097E0b23f64A496": newBigIntFromString("500000000000000000000"),
			},
		}},
	})
}

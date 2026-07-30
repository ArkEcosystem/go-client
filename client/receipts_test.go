package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestReceiptsService_All(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/receipts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/receipts?page=1&limit=1",
			    "first": "/api/receipts?page=1&limit=1",
			    "last": "/api/receipts?page=1&limit=1"
			  },
			  "data": [
			    {
			      "contractAddress": null,
			      "gasRefunded": 0,
			      "gasUsed": 21000,
			      "logs": [],
			      "output": "0x",
			      "status": 1,
			      "transactionHash": "178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c"
			    }
			  ]
			}`)
	})

	query := &ReceiptsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Receipts.All(context.Background(), query)
	testGeneralError(t, "Receipts.All", err)
	testResponseUrl(t, "Receipts.All", response, "/api/receipts")
	testResponseStruct(t, "Receipts.All", responseStruct, &Receipts{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/receipts?page=1&limit=1",
			First:      "/api/receipts?page=1&limit=1",
			Last:       "/api/receipts?page=1&limit=1",
		},
		Data: []Receipt{{
			ContractAddress: nil,
			GasRefunded:     0,
			GasUsed:         21000,
			Logs:            []ReceiptLog{},
			Output:          "0x",
			Status:          1,
			TransactionHash: "178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c",
		}},
	})
}

func TestReceiptsService_Get(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/receipts/178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "data": {
			    "contractAddress": null,
			    "gasRefunded": 0,
			    "gasUsed": 21000,
			    "logs": [],
			    "output": "0x",
			    "status": 1,
			    "transactionHash": "178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c"
			  }
			}`)
	})

	responseStruct, response, err := client.Receipts.Get(context.Background(), "178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c", nil)
	testGeneralError(t, "Receipts.Get", err)
	testResponseUrl(t, "Receipts.Get", response, "/api/receipts/178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c")
	testResponseStruct(t, "Receipts.Get", responseStruct, &GetReceipt{
		Data: Receipt{
			ContractAddress: nil,
			GasRefunded:     0,
			GasUsed:         21000,
			Logs:            []ReceiptLog{},
			Output:          "0x",
			Status:          1,
			TransactionHash: "178df6719bd55a792c0c935bc5cffcabb8a49532cf16ea97490a17114eb39a3c",
		},
	})
}

func TestReceiptsService_Contracts(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/receipts/contracts", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		fmt.Fprint(writer,
			`{
			  "meta": {
			    "count": 1,
			    "pageCount": 1,
			    "totalCount": 1,
			    "next": null,
			    "previous": null,
			    "self": "/api/receipts/contracts?page=1&limit=1",
			    "first": "/api/receipts/contracts?page=1&limit=1",
			    "last": "/api/receipts/contracts?page=1&limit=1"
			  },
			  "data": [
			    {
			      "contractAddress": "0x2A4ea729eD03B237a7411967832a7fA3a1A4ccb0",
			      "gasRefunded": 0,
			      "gasUsed": 385147,
			      "status": 1,
			      "transactionHash": "bbed674207435cf37bcd155f14e1eaef8165e2a55479b8d2a83c76b4b01d1509"
			    }
			  ]
			}`)
	})

	query := &ReceiptContractsQuery{Pagination: Pagination{Limit: 1}}
	responseStruct, response, err := client.Receipts.Contracts(context.Background(), query)
	testGeneralError(t, "Receipts.Contracts", err)
	testResponseUrl(t, "Receipts.Contracts", response, "/api/receipts/contracts")

	contractAddress := "0x2A4ea729eD03B237a7411967832a7fA3a1A4ccb0"
	testResponseStruct(t, "Receipts.Contracts", responseStruct, &Receipts{
		Meta: Meta{
			Count:      1,
			PageCount:  1,
			TotalCount: 1,
			Next:       nil,
			Previous:   nil,
			Self:       "/api/receipts/contracts?page=1&limit=1",
			First:      "/api/receipts/contracts?page=1&limit=1",
			Last:       "/api/receipts/contracts?page=1&limit=1",
		},
		Data: []Receipt{{
			ContractAddress: &contractAddress,
			GasRefunded:     0,
			GasUsed:         385147,
			Status:          1,
			TransactionHash: "bbed674207435cf37bcd155f14e1eaef8165e2a55479b8d2a83c76b4b01d1509",
		}},
	})
}

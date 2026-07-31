package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestEVMService_Call(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "POST")
		fmt.Fprint(writer,
			`{
			  "jsonrpc": "2.0",
			  "id": 1,
			  "result": "0x4"
			}`)
	})

	id := 1
	req := &EvmRequest{
		Method: "eth_blockNumber",
		Params: []interface{}{},
		Id:     &id,
	}
	responseStruct, response, err := client.EVM.Call(context.Background(), req)
	testGeneralError(t, "EVM.Call", err)
	testResponseUrl(t, "EVM.Call", response, "/evm/api/")
	testResponseStruct(t, "EVM.Call", responseStruct, map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      float64(1),
		"result":  "0x4",
	})

	if req.Jsonrpc != "2.0" {
		t.Errorf("EVM.Call did not default Jsonrpc to \"2.0\", got %q", req.Jsonrpc)
	}
}

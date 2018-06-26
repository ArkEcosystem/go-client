package main

import (
	"./ark_client"
	"./ark_client/structs/response"
	"context"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client := ark_client.NewClient(nil)

	response := &response.PublicKeyResponse{}

	accounts, _, _ := client.One_Accounts.PublicKey(context.Background(), "DQ7VAW7u171hwDW75R1BqfHbA9yiKRCBSh", &response)

	spew.Dump(accounts)
}

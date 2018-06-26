package main

import (
	"./ark_client"
	"context"
	"github.com/davecgh/go-spew/spew"
)

type PublicKeyResponse struct {
	Success   bool   `json:"success,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

func main() {
	client := ark_client.NewClient(nil)

	response := &PublicKeyResponse{}

	accounts, _, _ := client.One_Accounts.PublicKey(context.Background(), "DQ7VAW7u171hwDW75R1BqfHbA9yiKRCBSh", &response)

	spew.Dump(accounts)
}

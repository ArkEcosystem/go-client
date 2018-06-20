package main

import (
	"./ark_client"
	"context"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client := ark_client.NewClient(nil)

	// accounts, _, _ := client.One_Accounts.Top(context.Background())
	accounts, _, _ := client.One_Accounts.PublicKey(context.Background(), "DQ7VAW7u171hwDW75R1BqfHbA9yiKRCBSh")

	spew.Dump(accounts)
}

package main

import (
	"./client"
	"context"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client := client.NewClient(nil)

	accounts, _, _ := client.One_Accounts.PublicKey(context.Background(), "DQ7VAW7u171hwDW75R1BqfHbA9yiKRCBSh")

	spew.Dump(accounts.Success)
}

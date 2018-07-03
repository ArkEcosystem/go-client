package main

import (
	"context"
	"github.com/ArkEcosystem/go-client/client/one"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client := one.NewClient(nil)

	accounts, _, _ := client.Accounts.PublicKey(context.Background(), "DQ7VAW7u171hwDW75R1BqfHbA9yiKRCBSh")

	spew.Dump(accounts.Success)
}

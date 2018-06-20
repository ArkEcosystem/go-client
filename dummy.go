package main

import (
    "context"
    "./ark_client"
    "github.com/davecgh/go-spew/spew"
)

func main () {
    client := ark_client.NewClient(nil)

    resp, _ := client.One_Accounts.Top(context.Background())

    spew.Dump(resp)
}

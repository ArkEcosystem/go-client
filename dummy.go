package main

import (
    "context"
    "./ark_client"
)

func main () {
    client := ark_client.NewClient(nil)

    resp, _ := client.One_Accounts.List(context.Background())

    println(resp)
}

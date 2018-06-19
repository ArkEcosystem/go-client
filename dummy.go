package main

import (
    "context"
    "fmt"
    "./ark_client"
)

func main () {
    client := ark_client.NewClient(nil)

    resp, err := client.One_Accounts.List(context.Background())

    fmt.Println(resp, err)
}

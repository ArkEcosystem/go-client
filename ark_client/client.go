// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "bytes"
    // "context"
    "encoding/json"
    // "errors"
    "fmt"
    "io"
    // "io/ioutil"
    "net/http"
    "net/url"
    // "reflect"
    // "strconv"
    "strings"
    // "time"

    // "github.com/google/go-querystring/query"

    . "./types"
    "./one"
    "./two"
)

const (
    defaultBaseURL = "https://api.ark.io/"
    userAgent      = "go-client"
)

type Client struct {
    common Service

    // Services - Version 1
    One_Accounts *one.AccountsService
    One_Blocks *one.BlocksService
    One_Delegates *one.DelegatesService
    One_Loader *one.LoaderService
    One_Peers *one.PeersService
    One_Signatures *one.SignaturesService
    One_Transactions *one.TransactionsService

    // // Services - Version 2
    Two_Blocks *two.BlocksService
    Two_Delegates *two.DelegatesService
    Two_Node *two.NodeService
    Two_Peers *two.PeersService
    Two_Transactions *two.TransactionsService
    Two_Votes *two.VotesService
    Two_Wallets *two.WalletsService
    Two_Webhooks *two.WebhooksService
}

func NewClient(httpClient *http.Client) *Client {
    if httpClient == nil {
        httpClient = http.DefaultClient
    }
    baseURL, _ := url.Parse(defaultBaseURL)

    c := &Client{}
    c.common = Service {
        Client: &HttpClient {
            Client: httpClient,
            BaseURL: baseURL,
        },
    }

    c.One_Accounts = (*one.AccountsService)(&c.common)
    c.One_Blocks = (*one.BlocksService)(&c.common)
    c.One_Delegates = (*one.DelegatesService)(&c.common)
    c.One_Loader = (*one.LoaderService)(&c.common)
    c.One_Peers = (*one.PeersService)(&c.common)
    c.One_Signatures = (*one.SignaturesService)(&c.common)
    c.One_Transactions = (*one.TransactionsService)(&c.common)

    c.Two_Blocks = (*two.BlocksService)(&c.common)
    c.Two_Delegates = (*two.DelegatesService)(&c.common)
    c.Two_Node = (*two.NodeService)(&c.common)
    c.Two_Peers = (*two.PeersService)(&c.common)
    c.Two_Transactions = (*two.TransactionsService)(&c.common)
    c.Two_Votes = (*two.VotesService)(&c.common)
    c.Two_Wallets = (*two.WalletsService)(&c.common)
    c.Two_Webhooks = (*two.WebhooksService)(&c.common)

    return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
    if !strings.HasSuffix(c.common.Client.BaseURL.Path, "/") {
        return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.common.Client.BaseURL)
    }
    u, err := c.common.Client.BaseURL.Parse(urlStr)
    if err != nil {
        return nil, err
    }

    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        enc := json.NewEncoder(buf)
        enc.SetEscapeHTML(false)
        err := enc.Encode(body)
        if err != nil {
            return nil, err
        }
    }

    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }

    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }

    return req, nil
}

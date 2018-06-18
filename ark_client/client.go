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
    // "io"
    // "io/ioutil"
    "net/http"
    "net/url"
    // "reflect"
    // "strconv"
    "strings"
    "sync"
    // "time"

    // "github.com/google/go-querystring/query"

    "./one"
    "./two"
)

const (
    defaultBaseURL = "https://api.ark.io/"
    userAgent      = "go-client"
)

type Client struct {
    clientMu sync.Mutex
    client   *http.Client

    BaseURL *url.URL

    common service

    // Services - Version 1
    one_Accounts *one.AccountsService
    // one_Blocks *one.BlocksService
    // one_Delegates *one.DelegatesService
    // one_Loader *one.LoaderService
    // one_Peers *one.PeersService
    // one_Signatures *one.SignaturesService
    // one_Transactions *one.TransactionsService

    // // Services - Version 2
    // two_Blocks *two.BlocksService
    // two_Delegates *two.DelegatesService
    // two_Node *two.NodeService
    // two_Peers *two.PeersService
    // two_Transactions *two.TransactionsService
    // two_Votes *two.VotesService
    // two_Wallets *two.WalletsService
    // two_Webhooks *two.WebhooksService
}

type service struct {
    client *Client
}

func NewClient(httpClient *http.Client) *Client {
    if httpClient == nil {
        httpClient = http.DefaultClient
    }
    baseURL, _ := url.Parse(defaultBaseURL)

    c := &Client{client: httpClient, BaseURL: baseURL}
    c.common.client = c

    c.one.Accounts (*one_AccountsService)(&c.common)
    c.one.Blocks (*one_BlocksService)(&c.common)
    c.one.Delegates (*one_DelegatesService)(&c.common)
    c.one.Loader (*one_LoaderService)(&c.common)
    c.one.Peers (*one_PeersService)(&c.common)
    c.one.Signatures (*one_SignaturesService)(&c.common)
    c.one.Transactions (*one_TransactionsService)(&c.common)

    c.two.Blocks (*two_BlocksService)(&c.common)
    c.two.Delegates (*two_DelegatesService)(&c.common)
    c.two.Node (*two_NodeService)(&c.common)
    c.two.Peers (*two_PeersService)(&c.common)
    c.two.Transactions (*two_TransactionsService)(&c.common)
    c.two.Votes (*two_VotesService)(&c.common)
    c.two.Wallets (*two_WalletsService)(&c.common)
    c.two.Webhooks (*two_WebhooksService)(&c.common)

    return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
    if !strings.HasSuffix(c.BaseURL.Path, "/") {
        return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
    }
    u, err := c.BaseURL.Parse(urlStr)
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

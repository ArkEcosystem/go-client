// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "reflect"
    "strconv"
    "strings"
    "sync"
    "time"

    "github.com/google/go-querystring/query"
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
    v1_Accounts *AccountsService
    v1_Blocks *BlocksService
    v1_Delegates *DelegatesService
    v1_Loader *LoaderService
    v1_Peers *PeersService
    v1_Signatures *SignaturesService
    v1_Transactions *TransactionsService

    // Services - Version 2
    v2_Blocks *BlocksService
    v2_Delegates *DelegatesService
    v2_Node *NodeService
    v2_Peers *PeersService
    v2_Transactions *TransactionsService
    v2_Votes *VotesService
    v2_Wallets *WalletsService
    v2_Webhooks *WebhooksService
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

    c.one.Accounts (*v1_AccountsService)(&c.common)
    c.one.Blocks (*v1_BlocksService)(&c.common)
    c.one.Delegates (*v1_DelegatesService)(&c.common)
    c.one.Loader (*v1_LoaderService)(&c.common)
    c.one.Peers (*v1_PeersService)(&c.common)
    c.one.Signatures (*v1_SignaturesService)(&c.common)
    c.one.Transactions (*v1_TransactionsService)(&c.common)

    c.two.Blocks (*v2_BlocksService)(&c.common)
    c.two.Delegates (*v2_DelegatesService)(&c.common)
    c.two.Node (*v2_NodeService)(&c.common)
    c.two.Peers (*v2_PeersService)(&c.common)
    c.two.Transactions (*v2_TransactionsService)(&c.common)
    c.two.Votes (*v2_VotesService)(&c.common)
    c.two.Wallets (*v2_WalletsService)(&c.common)
    c.two.Webhooks (*v2_WebhooksService)(&c.common)

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

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArkEcosystem/go-client/client/one"
	"github.com/ArkEcosystem/go-client/client/two"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

const (
	defaultBaseURL = "https://dexplorer.ark.io:8443/api/"
	userAgent      = "go-client"
)

type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	BaseURL *url.URL

	common Service

	One_Accounts     *one.AccountsService
	One_Blocks       *one.BlocksService
	One_Delegates    *one.DelegatesService
	One_Loader       *one.LoaderService
	One_Peers        *one.PeersService
	One_Signatures   *one.SignaturesService
	One_Transactions *one.TransactionsService

	Two_Blocks       *two.BlocksService
	Two_Delegates    *two.DelegatesService
	Two_Node         *two.NodeService
	Two_Peers        *two.PeersService
	Two_Transactions *two.TransactionsService
	Two_Votes        *two.VotesService
	Two_Wallets      *two.WalletsService
}

type Service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c

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

	return c
}

func (c *Client) SendRequest(ctx context.Context, version int, method string, urlStr string, body interface{}, model interface{}) (*http.Response, error) {
	// Create a new HTTP request
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		if method == "POST" {
			buf = new(bytes.Buffer)
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			err := enc.Encode(body)
			if err != nil {
				return nil, err
			}
		}
	}

	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	if body != nil {
		if method == "GET" {
			params, _ := query.Values(body)

			req.URL.RawQuery = params.Encode()
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", strconv.Itoa(version))

	// Execute the previously created HTTP request
	resp, err := c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}

	// Map the JSON response to a struct
	if model != nil {
		if w, ok := model.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(model)

			if decErr == io.EOF {
				decErr = nil
			}

			if decErr != nil {
				err = decErr
			}
		}
	}

	defer resp.Body.Close()

	return resp, nil
}

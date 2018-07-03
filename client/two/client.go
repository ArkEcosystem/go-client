// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	Blocks       *BlocksService
	Delegates    *DelegatesService
	Node         *NodeService
	Peers        *PeersService
	Transactions *TransactionsService
	Votes        *VotesService
	Wallets      *WalletsService
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

	c.Blocks = (*BlocksService)(&c.common)
	c.Delegates = (*DelegatesService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Votes = (*VotesService)(&c.common)
	c.Wallets = (*WalletsService)(&c.common)

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
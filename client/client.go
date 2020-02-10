// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://dexplorer.ark.io:8443/api/"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	common Service

	Blocks       *BlocksService
	Bridgechains *BridgechainsService
	Businesses   *BusinessesService
	Delegates    *DelegatesService
	Locks        *LocksService
	Node         *NodeService
	Peers        *PeersService
	Rounds       *RoundsService
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
	c.Bridgechains = (*BridgechainsService)(&c.common)
	c.Businesses = (*BusinessesService)(&c.common)
	c.Delegates = (*DelegatesService)(&c.common)
	c.Locks = (*LocksService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Rounds = (*RoundsService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Votes = (*VotesService)(&c.common)
	c.Wallets = (*WalletsService)(&c.common)

	return c
}

func (c *Client) SendRequest(ctx context.Context, method string, urlStr string, queryString interface{}, body interface{}, model interface{}) (*http.Response, error) {
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
			json, _ := json.Marshal(body)
			buf = bytes.NewBuffer(json)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if queryString != nil {
		switch v := queryString.(type) {
		case *Pagination:
			if v.Page == 0 {
				v.Page = 1
			}
		}

		params, _ := query.Values(queryString)

		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Content-Type", "application/json")

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
	if w, ok := model.(io.Writer); ok {
		_, _ = io.Copy(w, resp.Body)
	} else {
		decErr := json.NewDecoder(resp.Body).Decode(model)

		if decErr == io.EOF {
			decErr = nil
		}

		if decErr != nil {
			err = decErr
		}
	}

	defer resp.Body.Close()

	return resp, err
}

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

type Hosts struct {
	API          string
	Transactions string
	EVM          string
}

type Client struct {
	httpClient *http.Client
	Hosts      Hosts

	common Service

	ApiNodes     *ApiNodesService
	Blocks       *BlocksService
	Blockchain   *BlockchainService
	Commits      *CommitsService
	Validators   *ValidatorsService
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

func NewClient(httpClient *http.Client, hosts Hosts) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
		Hosts:      hosts,
	}
	c.common.client = c

	c.ApiNodes = (*ApiNodesService)(&c.common)
	c.Blocks = (*BlocksService)(&c.common)
	c.Blockchain = (*BlockchainService)(&c.common)
	c.Commits = (*CommitsService)(&c.common)
	c.Validators = (*ValidatorsService)(&c.common)
	c.Node = (*NodeService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Rounds = (*RoundsService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Votes = (*VotesService)(&c.common)
	c.Wallets = (*WalletsService)(&c.common)

	return c
}

func (c *Client) SendRequest(ctx context.Context, method string, endpoint string, queryString interface{}, body interface{}, model interface{}, hostType string) (*http.Response, error) {
	var host string
	switch hostType {
	case "transactions":
		host = c.Hosts.Transactions
	case "evm":
		host = c.Hosts.EVM
	default:
		host = c.Hosts.API
	}

	parsedHost, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("invalid host URL: %v", err)
	}

	if !strings.HasSuffix(parsedHost.Path, "/") {
		parsedHost.Path += "/"
	}

	u, err := parsedHost.Parse(endpoint)
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
	resp, err := c.httpClient.Do(req)

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

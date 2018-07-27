// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://dexplorer.ark.io:8443/api/"
)

type Client struct {
	client   *http.Client

	BaseURL *url.URL

	common Service

	Accounts     *AccountsService
	Blocks       *BlocksService
	Delegates    *DelegatesService
	Loader       *LoaderService
	Peers        *PeersService
	Signatures   *SignaturesService
	Transactions *TransactionsService
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

	c.Accounts = (*AccountsService)(&c.common)
	c.Blocks = (*BlocksService)(&c.common)
	c.Delegates = (*DelegatesService)(&c.common)
	c.Loader = (*LoaderService)(&c.common)
	c.Peers = (*PeersService)(&c.common)
	c.Signatures = (*SignaturesService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)

	return c
}

func (c *Client) SendRequest(ctx context.Context, method string, urlStr string, body interface{}, model interface{}) (*http.Response, error) {
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
	req.Header.Set("API-Version", "1")

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
	}

	defer resp.Body.Close()

	return resp, err
}

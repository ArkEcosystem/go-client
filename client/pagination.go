// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"net/url"
	"strings"
)

type Pagination struct {
	Page   int `url:"page"`
	Limit  int `url:"limit"`
	Offset int `url:"offset,omitempty"`
}

// paginator is implemented by *Pagination and by any query type that embeds
// Pagination (method promotion carries the implementation through). This
// lets SendRequest apply the "default page to 1" behavior uniformly, since
// a plain type switch on *Pagination wouldn't match embedding types.
type paginator interface {
	applyDefaults()
}

func (p *Pagination) applyDefaults() {
	if p == nil {
		return
	}

	if p.Page == 0 {
		p.Page = 1
	}
}

// CommaSeparated encodes a slice of strings as a single comma-joined query
// value (e.g. addresses=a,b,c) instead of go-querystring's default of
// repeating the key for each value.
type CommaSeparated []string

func (c CommaSeparated) EncodeValues(key string, v *url.Values) error {
	if len(c) == 0 {
		return nil
	}

	v.Set(key, strings.Join(c, ","))

	return nil
}

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
	Page  int `url:"page"`
	Limit int `url:"limit"`
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

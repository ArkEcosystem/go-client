// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Meta struct {
	Count      uint32 `url:"count,omitempty"`
	PageCount  uint32 `url:"pageCount,omitempty"`
	TotalCount uint32 `url:"totalCount,omitempty"`
	Next       string `url:"next,omitempty"`
	Previous   string `url:"previous,omitempty"`
	Self       string `url:"self,omitempty"`
	First      string `url:"first,omitempty"`
	Last       string `url:"last,omitempty"`
}

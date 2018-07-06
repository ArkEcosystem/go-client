// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

type Meta struct {
	Count      byte   `url:"count,omitempty"`
	PageCount  byte   `url:"pageCount,omitempty"`
	TotalCount byte   `url:"totalCount,omitempty"`
	Next       string `url:"next,omitempty"`
	Previous   string `url:"previous,omitempty"`
	Self       string `url:"self,omitempty"`
	First      string `url:"first,omitempty"`
	Last       string `url:"last,omitempty"`
}

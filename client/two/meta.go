// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

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

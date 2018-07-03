// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package request

type AddressQuery struct {
	Address string `url:"address"`
}

type TopQuery struct {
	Limit  int `url:"limit"`
	Offset int `url:"offset"`
}

// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package requests_one

type GetPeersQuery struct {
	Port    int    `url:"port"`
	Status  string `url:"status"`
	Os      string `url:"os"`
	Version string `url:"version"`
	OrderBy string `url:"orderBy"`
	Limit   int    `url:"limit"`
	Offset  int    `url:"offset"`
}

type GetPeerQuery struct {
	Ip   string `url:"ip"`
	Port int    `url:"port"`
}

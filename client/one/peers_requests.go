// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type GetPeersQuery struct {
	Port    byte   `url:"port"`
	Status  string `url:"status"`
	Os      string `url:"os"`
	Version string `url:"version"`
	OrderBy string `url:"orderBy"`
	Limit   byte   `url:"limit"`
	Offset  byte   `url:"offset"`
}

type GetPeerQuery struct {
	Ip   string `url:"ip"`
	Port byte   `url:"port"`
}

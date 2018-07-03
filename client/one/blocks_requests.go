// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type BlocksQuery struct {
	Limit              int    `url:"limit"`
	OrderBy            string `url:"orderBy"`
	Offset             int    `url:"offset"`
	GeneratorPublicKey string `url:"generatorPublicKey"`
	TotalAmount        int    `url:"totalAmount"`
	TotalFee           int    `url:"totalFee"`
	Reward             int    `url:"reward"`
	PreviousBlock      string `url:"previousBlock"`
	Height             int    `url:"height"`
}

type BlockIdQuery struct {
	Id string `url:"id"`
}

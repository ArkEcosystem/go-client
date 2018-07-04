// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type BlocksQuery struct {
	Limit              byte   `url:"limit"`
	OrderBy            string `url:"orderBy"`
	Offset             byte   `url:"offset"`
	GeneratorPublicKey string `url:"generatorPublicKey"`
	TotalAmount        byte   `url:"totalAmount"`
	TotalFee           byte   `url:"totalFee"`
	Reward             byte   `url:"reward"`
	PreviousBlock      string `url:"previousBlock"`
	Height             byte   `url:"height"`
}

type BlockIdQuery struct {
	Id string `url:"id"`
}

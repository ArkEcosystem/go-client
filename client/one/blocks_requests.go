// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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

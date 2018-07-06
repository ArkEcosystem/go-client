// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

type GetPeersQuery struct {
	Port    uint16 `url:"port"`
	Status  string `url:"status"`
	Os      string `url:"os"`
	Version string `url:"version"`
	OrderBy string `url:"orderBy"`
	Limit   byte   `url:"limit"`
	Offset  byte   `url:"offset"`
}

type GetPeerQuery struct {
	Ip   string `url:"ip"`
	Port uint16 `url:"port"`
}

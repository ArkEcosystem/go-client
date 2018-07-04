// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

type GeneratorPublicKeyQuery struct {
	GeneratorPublicKey string `url:"generatorPublicKey"`
}

type PublicKeyQuery struct {
	PublicKey string `url:"publicKey"`
}

type GetDelegateQuery struct {
	PublicKey string `url:"publicKey"`
	Username  string `url:"username"`
}

type GetDelegatesQuery struct {
	OrderBy string `url:"orderBy"`
	Limit   byte   `url:"limit"`
	Offset  byte   `url:"offset"`
}

type DelegateSearchQuery struct {
	Q     string `url:"q"`
	Limit byte   `url:"limit"`
}

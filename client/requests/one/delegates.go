// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package requests_one

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
	Limit   int    `url:"limit"`
	Offset  int    `url:"offset"`
}

type DelegateSearchQuery struct {
	Q     string `url:"q"`
	Limit string `url:"limit"`
}

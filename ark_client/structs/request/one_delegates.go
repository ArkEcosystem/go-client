// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package request

type GeneratorPublicKeyQuery struct {
	GeneratorPublicKey string `url:"generatorPublicKey"`
}

type PublicKeyQuery struct {
	PublicKey string `url:"publicKey"`
}

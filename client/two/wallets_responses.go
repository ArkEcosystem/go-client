// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package two

type Wallet struct {
	Address    string `json:"address,omitempty"`
	Address    string `json:"address,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
	Balance    int64  `json:"balance,omitempty"`
	IsDelegate bool   `json:"isDelegate,omitempty"`
}

type Wallets struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Wallet `json:"data,omitempty"`
}

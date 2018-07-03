// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type PublicKey struct {
	Success   bool   `json:"success,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type Account struct {
	Address                     string    `json:"address,omitempty"`
	PublicKey                   string    `json:"publicKey,omitempty"`
	SecondPublicKey             string    `json:"secondPublicKey,omitempty"`
	Username                    string    `json:"username,omitempty"`
	Balance                     string    `json:"balance,omitempty"`
	UnconfirmedBalance          string    `json:"unconfirmedBalance,omitempty"`
	Multisignatures             []string  `json:"multisignatures,omitempty"`
	UnconfirmedMultisignatures  []string  `json:"u_multisignatures,omitempty"`
	UnconfirmedSignature        int       `json:"unconfirmedSignature,omitempty"`
	SecondSignature             int       `json:"secondSignature,omitempty"`
}

type Accounts struct {
	Success   bool       `json:"success,omitempty"`
	Accounts  []Account  `json:"accounts,omitempty"`
}

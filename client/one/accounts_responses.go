// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type Account struct {
	Address                    string   `json:"address,omitempty"`
	PublicKey                  string   `json:"publicKey,omitempty"`
	SecondPublicKey            string   `json:"secondPublicKey,omitempty"`
	Username                   string   `json:"username,omitempty"`
	Balance                    string   `json:"balance,omitempty"`
	UnconfirmedBalance         string   `json:"unconfirmedBalance,omitempty"`
	Multisignatures            []string `json:"multisignatures,omitempty"`
	UnconfirmedMultisignatures []string `json:"u_multisignatures,omitempty"`
	UnconfirmedSignature       int      `json:"unconfirmedSignature,omitempty"`
	SecondSignature            int      `json:"secondSignature,omitempty"`
}

type AccountSingle struct {
	Success  bool    `json:"success,omitempty"`
	Account  Account `json:"account,omitempty"`
}

type Accounts struct {
	Success  bool      `json:"success,omitempty"`
	Accounts []Account `json:"accounts,omitempty"`
}

type AccountsCount struct {
	Success bool `json:"success,omitempty"`
	Count   int  `json:"count,omitempty"`
}

type AccountDelegates struct {
	Success   bool       `json:"success,omitempty"`
	Delegates []Delegate `json:"delegates,omitempty"`
}

type AccountBalance struct {
	Success            bool   `json:"success,omitempty"`
	Balance            string `json:"balance,omitempty"`
	UnconfirmedBalance string `json:"unconfirmedBalance,omitempty"`
}

type AccountsTop struct {
	Success  bool      `json:"success,omitempty"`
	Accounts []Account `json:"accounts,omitempty"`
}

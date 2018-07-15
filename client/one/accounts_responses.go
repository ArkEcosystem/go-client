// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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
	UnconfirmedSignature       byte     `json:"unconfirmedSignature,omitempty"`
	SecondSignature            byte     `json:"secondSignature,omitempty"`
}

type AccountSingle struct {
	Success bool    `json:"success,omitempty"`
	Account Account `json:"account,omitempty"`
}

type Accounts struct {
	Success  bool      `json:"success,omitempty"`
	Accounts []Account `json:"accounts,omitempty"`
}

type AccountsCount struct {
	Success bool `json:"success,omitempty"`
	Count   byte `json:"count,omitempty"`
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

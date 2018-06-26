// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package response

type PublicKeyResponse struct {
	Success   bool   `json:"success,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

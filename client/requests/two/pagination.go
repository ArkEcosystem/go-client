// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package requests_two

type Pagination struct {
	Page  int `url:"page"`
	Limit int `url:"limit"`
}

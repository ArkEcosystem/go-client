// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type LoaderStatus struct {
	Success     bool  `json:"success,omitempty"`
	Loaded      bool  `json:"loaded,omitempty"`
	Now         int64 `json:"now,omitempty"`
	BlocksCount byte  `json:"blocksCount,omitempty"`
}

type LoaderSync struct {
	Success bool   `json:"success,omitempty"`
	Syncing bool   `json:"syncing,omitempty"`
	Blocks  byte   `json:"blocks,omitempty"`
	Height  int64  `json:"height,omitempty"`
	Id      string `json:"id,omitempty"`
}

type LoaderAutoConfigure struct {
	Success       bool            `json:"success,omitempty"`
	Nethash       string          `json:"nethash,omitempty"`
	Token         string          `json:"token,omitempty"`
	Symbol        string          `json:"symbol,omitempty"`
	Explorer      string          `json:"explorer,omitempty"`
	Version       byte            `json:"version,omitempty"`
	Ports         map[string]byte `json:"ports,omitempty"`
	FeeStatistics []FeeStatistic  `json:"feeStatistics,omitempty"`
}

type FeeStatistic struct {
	Type int64 `json:"type,omitempty"`
	Fees struct {
		MinFee int64 `json:"minFee,omitempty"`
		MaxFee int64 `json:"maxFee,omitempty"`
		MvgFee int64 `json:"avgFee,omitempty"`
	} `json:"fees,omitempty"`
}

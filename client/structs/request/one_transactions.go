// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package request

type TransactionIdQuery struct {
	Id string `url:"id"`
}

type GetUnconfirmedTransactionsQuery struct {
	SenderPublicKey string `url:"senderPublicKey"`
}

type GetTransactionsQuery struct {
	BlockId         string `url:"blockId"`
	Limit           int    `url:"limit"`
	Type            int    `url:"type"`
	OrderBy         string `url:"orderBy"`
	Offset          int    `url:"offset"`
	SenderPublicKey string `url:"senderPublicKey"`
	VendorField     string `url:"vendorField"`
	OwnerPublicKey  string `url:"ownerPublicKey"`
	OwnerAddress    string `url:"ownerAddress"`
	SenderId        string `url:"senderId"`
	RecipientId     string `url:"recipientId"`
	Amount          int    `url:"amount"`
	Fee             int    `url:"fee"`
}

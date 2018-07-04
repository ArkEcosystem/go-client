// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package one

type TransactionIdQuery struct {
	Id string `url:"id"`
}

type GetUnconfirmedTransactionsQuery struct {
	SenderPublicKey string `url:"senderPublicKey"`
}

type GetTransactionsQuery struct {
	BlockId         string `url:"blockId"`
	Limit           byte   `url:"limit"`
	Type            byte   `url:"type"`
	OrderBy         string `url:"orderBy"`
	Offset          byte   `url:"offset"`
	SenderPublicKey string `url:"senderPublicKey"`
	VendorField     string `url:"vendorField"`
	OwnerPublicKey  string `url:"ownerPublicKey"`
	OwnerAddress    string `url:"ownerAddress"`
	SenderId        string `url:"senderId"`
	RecipientId     string `url:"recipientId"`
	Amount          byte   `url:"amount"`
	Fee             byte   `url:"fee"`
}

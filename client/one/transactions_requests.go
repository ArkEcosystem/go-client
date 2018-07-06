// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

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

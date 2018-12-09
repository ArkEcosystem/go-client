// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type WalletsSearchRequest struct {
	OrderBy         string  `json:"orderBy,omitempty"`
	Address         string  `json:"address,omitempty"`
	PublicKey       string  `json:"publicKey,omitempty"`
	SecondPublicKey string  `json:"secondPublicKey,omitempty"`
	Vote            string  `json:"vote,omitempty"`
	Username        string  `json:"username,omitempty"`
	ProducedBlocks  uint32  `json:"producedBlocks,omitempty"`
	MissedBlocks    uint32  `json:"missedBlocks,omitempty"`
	Balance         *FromTo `json:"balance ,omitempty"`
	VoteBalance     *FromTo `json:"voteBalance ,omitempty"`
}

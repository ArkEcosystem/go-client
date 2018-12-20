// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"encoding/json"
	"strconv"
)

type FromTo struct {
	From interface{} `json:"from,omitempty"`
	To   interface{} `json:"from,omitempty"`
}

type FlexToshi uint64

func (fi *FlexToshi) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*uint64)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*fi = FlexToshi(i)

	return nil
}

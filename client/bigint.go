package client

import (
	"fmt"
	"math/big"
	"strings"
)

type BigInt struct {
	*big.Int
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		b.Int = big.NewInt(0)
		return nil
	}

	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return fmt.Errorf("client: invalid big integer %q", s)
	}

	b.Int = i

	return nil
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte(`"0"`), nil
	}

	return []byte(`"` + b.Int.String() + `"`), nil
}

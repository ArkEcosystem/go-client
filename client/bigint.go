package client

import (
	"fmt"
	"math/big"
	"strings"
)

// BigInt wraps *big.Int to (un)marshal the decimal-string format the API
// uses for wei-scale values (nonces, balances, fees, rewards), avoiding the
// range ceiling of uint64/uint32.
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

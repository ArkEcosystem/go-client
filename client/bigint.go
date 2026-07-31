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

// UnmarshalJSON and MarshalJSON intentionally use different receiver types
// (a lint tool may flag this). UnmarshalJSON must be a pointer receiver since
// it mutates b.Int; a value receiver would only update a local copy and the
// parsed value would be lost. MarshalJSON must stay a value receiver because
// BigInt is used as a map value elsewhere (e.g. NodeFeesResponse), and map
// values are not addressable in Go — a pointer receiver would fail to satisfy
// json.Marshaler there, silently falling back to *big.Int's default (bare,
// unquoted) JSON encoding instead of this type's quoted decimal-string format.
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

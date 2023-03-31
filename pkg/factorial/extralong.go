package factorial

import (
	"math/big"
)

func Factorial(n int32) string {
	out := big.NewInt(int64(1))
	for x := n; x > 1; x-- {
		out = out.Mul(out, big.NewInt(int64(x)))
	}
	return out.String()
}

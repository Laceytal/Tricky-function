package functions

import "math/big"

func GFunc(n string) string {
	N := new(big.Int)
	N.SetString(n, 10)
	if N.Sign() == 0 {
		return "NaN"
	}
	f1 := reverseTrim(n)
	F1 := new(big.Int)
	F1.SetString(f1, 10)
	f2 := reverseTrim(f1)
	F2 := new(big.Int)
	F2.SetString(f2, 10)
	r := new(big.Rat).SetFrac(F2, N)
	return r.RatString()
}

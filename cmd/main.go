package main

import (
	"TrickyFunction/internal/functions"
	"fmt"
	"math/big"
)

func main() {
	unique := make(map[string]struct{})

	for length := 1; length <= 30; length++ {
		min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
		max := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil), big.NewInt(1))

		n := new(big.Int).Set(min)
		one := big.NewInt(1)

		for n.Cmp(max) <= 0 {
			s := n.String()
			val := functions.GFunc(s)
			unique[val] = struct{}{}
			n.Add(n, one)
		}
		fmt.Printf("Длина %d обработана\n", length)
	}

	fmt.Println("Итог:", len(unique))
}

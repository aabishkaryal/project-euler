package main

import (
	"fmt"
	"math/big"
)

func main() {
	// n := new(big.Int).Exp(big.NewInt(12), big.NewInt(123), big.NewInt(0))
	n := big.NewInt(600851475143)
	biggestPrimeFactor := biggestPrimeFactor(n)

	fmt.Println(biggestPrimeFactor)
}

func biggestPrimeFactor(n *big.Int) *big.Int {
	var (
		lastFactor = new(big.Int)
		factor     = big.NewInt(3)
	)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	for new(big.Int).Rem(n, two).Cmp(zero) == 0 {
		lastFactor.Set(two)
		n.Div(n, two)
	}
	maxFactor := new(big.Int).Sqrt(n)
	for n.Cmp(one) == 1 && factor.Cmp(maxFactor) == -1 {
		for new(big.Int).Rem(n, factor).Cmp(zero) == 0 {
			lastFactor.Set(factor)
			n.Div(n, factor)
		}
		factor.Add(factor, two)
	}
	if n.Cmp(one) == 0 {
		return lastFactor
	} else {
		return n
	}
}

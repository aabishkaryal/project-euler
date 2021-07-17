package main

import (
	"math"
	"math/big"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	tests := []struct {
		limit int
	}{
		{int(math.Pow(1000000000, 0.5))}, {1000000000},
	}
	for _, test := range tests {
		primes := GeneratePrimes(test.limit)
		for _, n := range primes {
			if !big.NewInt(int64(n)).ProbablyPrime(40) {
				t.Errorf("%d is not Prime\n", n)
			}
		}
	}
}

package main

import (
	"fmt"
	"math"
	"math/big"
)

var composite map[int]bool

func main() {
	limit := int(math.Pow10(9))
	separatePrimes(limit)
	for k, v := range composite {
		if !v {
			if !big.NewInt(int64(k)).ProbablyPrime(40) {
				panic(k)
			}
			fmt.Println(k)
		}
	}
}

func separatePrimes(limit int) {
	composite = make(map[int]bool)
	composite[0] = true
	composite[1] = true
	for i := 2; i < limit; i++ {
		if !composite[i] {
			composite[i] = false
			for j := i * i; j < limit; j += i {
				composite[j] = true
			}
		}
	}
}

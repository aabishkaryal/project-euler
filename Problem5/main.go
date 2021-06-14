package main

import (
	"fmt"
)

// var zero *big.Int = big.NewInt(0)

func main() {
	n := 20
	fmt.Println(SmallestMultiple(n))
}

func SmallestMultiple(n int) int {
	lcm := 1
	for i := 1; i <= n; i++ {
		lcm = LCM(lcm, i)
	}
	return lcm
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

package main

import (
	"fmt"
)

func main() {
	n := 101
	cache := GeneratePrimes(n)
	fmt.Println(cache[n-1])
}

func GeneratePrimes(n int) map[int]int {
	cache := make(map[int]int)
	counter := 0
	for i := 2; counter < n; i++ {
		isPrime := true
		for j := 0; j < counter; j++ {
			if i%cache[j] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			cache[counter] = i
			counter++
		}
	}
	return cache
}

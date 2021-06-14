package main

import "fmt"

func main() {
	n := 100
	fmt.Println(SumOfNumbers(n)*SumOfNumbers(n) - SumOfSquares(n))
}

func SumOfNumbers(n int) int {
	return n * (n + 1) / 2
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

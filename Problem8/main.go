package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numAdjacentDigits := 13
	num := readFile("num.txt")
	p := greatestProductWithAdjacentDigits(num, numAdjacentDigits)
	fmt.Printf("The greatest product with %d ajacent digits is: %d\n", numAdjacentDigits, p)
}
func greatestProductWithAdjacentDigits(num string, nDigits int) int {
	highestProduct := 0
	for i := 0; i <= len(num)-nDigits; i++ {
		product := digitProduct(num[i : i+nDigits])
		if product > highestProduct {
			highestProduct = product
		}
	}
	return highestProduct
}

func digitProduct(num string) int {
	product := 1
	for i := 0; i < len(num); i++ {
		n, err := strconv.Atoi(string(num[i]))
		if err != nil {
			panic(err)
		}
		product *= n
	}
	return product
}

func readFile(fileName string) string {
	num, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error occured while reading file %s\n", fileName)
		panic(err)
	}
	return string(num)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	digit := 5
	fmt.Println(BiggestPalindromeWithDigits(digit))
}

func BiggestPalindromeWithDigits(d int) (largestPalindrome int) {
	largestPalindrome = 0
	for i := int(math.Pow10(d) - 1); i > int(math.Pow10(d-1)-1); i-- {
		for j := i; j > int(math.Pow10(d-1)-1); j-- {
			n := i * j
			if n > largestPalindrome && isPalindrome(n) {
				largestPalindrome = n
			}
		}
	}
	return largestPalindrome
}

func isPalindrome(n int) bool {
	reverse := 0
	original := n
	for n > 0 {
		reverse = 10*reverse + n%10
		n = n / 10
	}
	return reverse == original
}

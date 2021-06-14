package main

import "testing"

var result int

func benchmarkBiggestPalindromeWithDigits(d int, b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = BiggestPalindromeWithDigits(d)
	}
	result = r
}

func BenchmarkBiggestPalindromeWithDigits2(b *testing.B) {
	benchmarkBiggestPalindromeWithDigits(2, b)
}

func BenchmarkBiggestPalindromeWithDigits3(b *testing.B) {
	benchmarkBiggestPalindromeWithDigits(3, b)
}

func BenchmarkBiggestPalindromeWithDigits4(b *testing.B) {
	benchmarkBiggestPalindromeWithDigits(4, b)
}

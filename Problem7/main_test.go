package main

import "testing"

var result map[int]int = make(map[int]int)

func benchmarkGeneratePrimes(n int, b *testing.B) {
	var r map[int]int = make(map[int]int)
	for i := 0; i < b.N; i++ {
		r = GeneratePrimes(n)
	}
	result = r
}

func BenchmarkGeneratePrimes100(b *testing.B) {
	benchmarkGeneratePrimes(100, b)
}

func BenchmarkGeneratePrimes1000(b *testing.B) {
	benchmarkGeneratePrimes(1000, b)
}

func BenchmarkGeneratePrimes10000(b *testing.B) {
	benchmarkGeneratePrimes(10000, b)
}

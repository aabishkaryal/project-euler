package main

import "testing"

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution(1000, []int{3, 5})
	}
}

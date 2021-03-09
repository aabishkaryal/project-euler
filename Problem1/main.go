package main

import (
	"fmt"
)

func main() {
	limit := 1000
	nums := []int{3, 5}
	sum := Solution(limit, nums)

	fmt.Println(sum)
}

func Solution(limit int, nums []int) (sum int) {
	for i := 0; i < limit; i++ {
		for _, n := range nums {
			if i%n == 0 {
				sum += i
				break
			}
		}
	}
	return
}


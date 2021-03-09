package main

import "fmt"

func main() {
	limit := 4000000
	sum := Solution(limit)

	fmt.Println(sum)
}

func Solution(limit int) (sum int) {
	current, prev := 1, 1

	for current < limit {
		current, prev = current+prev, current
		if current%2 == 0 {
			sum += current
		}
	}
	return
}

package main

import "fmt"

func main() {
	limit := 4000000
	sum := 0

	current, prev := 1, 1

	for current < limit {
		current, prev = current+prev, current
		if current%2 == 0 {
			sum += current
		}
	}

	fmt.Println(sum)
}

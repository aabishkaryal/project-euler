package main

import(
	"fmt"
)

func main(){
	limit := 1000
	nums := [2]int{3, 5}
	sum := 0
	for i := 0; i < limit; i++{
		for _, n := range nums{
			if i % n == 0{
				sum += i
				break
			}
		} 		
	}
	fmt.Println(sum)
}

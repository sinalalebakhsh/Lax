package main

import (
	"fmt"
)

func sum_numbers(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum_numbers(n-1)
}

func main()  {
	fmt.Println(sum_numbers(20))
}
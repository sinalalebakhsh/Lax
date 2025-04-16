package main

import "fmt"

func main() {
	numbers := make([]int, 0, 4)
	for i := 0; i < 6; i++ {
		fmt.Println(
			"length:",
			len(numbers), 
			"capacity:",
			cap(numbers),
			"numbers:",
			numbers,
		)
		numbers = append(numbers, i)
	}
}

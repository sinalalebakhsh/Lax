package main

import (
	"fmt"
)

func main()  {
	numbers := []int{1,2,3,4}
	copy_to_numbers := copy(numbers, []int{77, 99})
	fmt.Println(
		"numbers:",
		numbers,
		"copy_to_numbers:",
		copy_to_numbers,
	)
	
}
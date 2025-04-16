package main

import "fmt"

func main()  {
	numbers := make([]int, 2, 4)
	
	numbers = append(numbers, 1,2,3,4,5,6)

	fmt.Println(numbers)

	numbers[0] = 99
	numbers[1] = 66

	fmt.Println(numbers)

}
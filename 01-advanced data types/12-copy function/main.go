package main

import "fmt"

func main()  {
	numbers := make([]int,4)

	fmt.Println(numbers)


	// ===============================================
	another := []int{1,2,3,4}

	copy(numbers, another)

	fmt.Println(numbers)





	
}
package main

import "fmt"

func main()  {
	var numbers_0_to_100 [100]int


	for i := 0; i < len(numbers_0_to_100); i++ {
		// numbers_0_to_100[i] = i	
		numbers_0_to_100[i] = i+1	
	}
	
	fmt.Println(numbers_0_to_100)
}
package main

import "fmt"

func main()  {
	numbers := []int{1,2,3,4,5,6}

	for range numbers {
		fmt.Println("I Love Go")
	}
	
}
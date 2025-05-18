package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5}

	for index, number := range numbers {
		fmt.Println("index:",index, "number:", number)
	}

	
}
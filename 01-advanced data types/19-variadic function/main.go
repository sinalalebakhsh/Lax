package main

import "fmt"

func sum(numbers ...int) int  {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main()  {
	fmt.Println(sum(1,2,2,2,4))
	
	slice_numbers := []int{5,5,5,1}
	fmt.Println(sum(slice_numbers...))
}
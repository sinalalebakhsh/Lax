package main

import "fmt"

func main()  {
	var number int = 5
	var address_of_number_in_memory *int = &number
	fmt.Println("number:", number)
	fmt.Println("address of number in memory:", address_of_number_in_memory)

	var a_number *int 
	var b_number *int = a_number
	fmt.Println("number:", a_number)
	fmt.Println("number:", b_number)
	
}
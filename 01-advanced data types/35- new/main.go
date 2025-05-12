package main

import "fmt"


func main()  {
	// var number int = 5
	a_number := new(int)
	
	fmt.Println(a_number)
	
	// ---------------------------

	fmt.Println(*a_number)

	
	*a_number = 354
	fmt.Println(*a_number)

}
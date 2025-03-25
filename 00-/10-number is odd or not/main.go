package main

import "fmt"

func main()  {
	var A_number int = 5
	if A_number % 2 == 0 {
		fmt.Println(A_number, " is even.")
	} else {
		fmt.Println(A_number, "is odd.")
	}
	
}
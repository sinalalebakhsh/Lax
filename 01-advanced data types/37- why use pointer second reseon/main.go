package main

import "fmt"


func change_a_number(number *int)  {
	*number = 111
	fmt.Println(*number)
}


func main()  {
	var number int = 44
	fmt.Println(number)


	change_a_number(&number)

	fmt.Println(number)
}




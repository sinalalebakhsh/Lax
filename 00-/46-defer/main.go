package main

import "fmt"


func function_defer_()  {
	fmt.Println("1")
}


func main()  {
	defer fmt.Println("2")
	defer fmt.Println("3")
	function_defer_()
	defer fmt.Println("4")

	defer func()  {
		fmt.Println("defer funct()")
	}()

	defer fmt.Println("5")
}
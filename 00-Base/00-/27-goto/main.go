package main

import "fmt"

func main()  {
	goto b
	fmt.Println("a")
	b:
	fmt.Println("b")
	
}
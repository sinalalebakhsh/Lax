package main

import "fmt"


const (
	true_works = iota
	false_works
	equal_works
	another_works
)


func main()  {

	fmt.Println("Good works have", true_works , "code.")
	fmt.Println("Bad works have", false_works , "code.")
	
}
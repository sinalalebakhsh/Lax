package main

import "fmt"

func main() {
	a_number := 1
	next := func() int {
		a_number++
		return a_number
	}
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
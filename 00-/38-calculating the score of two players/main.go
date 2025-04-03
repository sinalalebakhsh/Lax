package main

import "fmt"

func main() {
	sina := 0
	lina := 0

	sina_score := func() int {
		sina++
		return sina
	}
	mina_score := func() int {
		lina++
		return lina
	}

	fmt.Println("sina score =",sina_score())
	fmt.Println(mina_score())
	fmt.Println(mina_score())
	fmt.Println(mina_score())
	fmt.Println("sina score =",sina_score())
	fmt.Println(mina_score())
	fmt.Println("sina score =",sina_score())
	fmt.Println("sina score =",sina_score())
	fmt.Println("sina score =",sina_score())

}
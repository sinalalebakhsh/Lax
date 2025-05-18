package main

import (
	"fmt"
)

func new_counter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main()  {
	sina := new_counter()
	lina := new_counter()
	dina := new_counter()

	fmt.Println("sina=", sina())
	fmt.Println("sina=", sina())
	fmt.Println("lina=", lina())
	fmt.Println("sina=", sina())
	fmt.Println("dina=", dina())
	fmt.Println("sina=", sina())
	fmt.Println("sina=", sina())
	

}
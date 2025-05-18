package main

import "fmt"

func main() {

	var a_slice []int
	fmt.Println("slice a_slice:",a_slice)
	fmt.Println("length:", len(a_slice))


	fmt.Println("==============================")
	new_a_slice := append(a_slice, 1)
	fmt.Println("slice new_a_slice:",new_a_slice)
	fmt.Println("length:", len(new_a_slice))

}

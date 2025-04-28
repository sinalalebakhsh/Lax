package main

import (
	"fmt"
	"reflect"
)

func main() {
	src := []int{10, 11, 12, 13, 14}
	dst := []int{0, 1, 2, 3, 4}
	n := copy(dst, src)
	fmt.Println("n =", n, "src =", src, "dst =", dst)
	
	
	
	

	numbers := []int{666,888, 555, 222, 12, 13, 13 ,15}
	destination_numbers := []int{1, 0, 1, 0}
	number_of_destination_numbers := copy(destination_numbers, numbers)
	fmt.Println("numbers:",numbers)
	fmt.Println("number_of_destination_numbers:", number_of_destination_numbers)
	fmt.Println("destination_numbers:",destination_numbers)
	
	

	fmt.Println(reflect.TypeOf(numbers))
	fmt.Println(reflect.TypeOf(destination_numbers))
	fmt.Println(reflect.TypeOf(number_of_destination_numbers))
}

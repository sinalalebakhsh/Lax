package main

import "fmt"

func double(number_ int) int {
	return number_ * 2
}

func square(number_ int) int {
	return number_ * number_
}

func mohite_moraba(zele_moraba int) int {
	return zele_moraba * 4
}


func apply_function(numbers int, functions func(int) int) {
	result := functions(numbers)
	fmt.Println(result)
}

func main() {
	apply_function(2 , double)   

	apply_function(3, square)

	apply_function(5, mohite_moraba)


}

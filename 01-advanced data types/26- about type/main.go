package main

import "fmt"

type number int

func add_numbers(a, b number) number {
	return a + b
}

func main()  {

	var number_one number = 1
	var number_two number = 2
	
	sum_numbers := add_numbers(number_one, number_two)


	fmt.Println(sum_numbers)

}
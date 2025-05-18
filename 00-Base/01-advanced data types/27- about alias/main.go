package main

import "fmt"


// alias is it a equal sign =
type number = int

func add_numbers(a, b number) number {
	return a + b
}

func main()  {

	// differences is here  
	number_one := 1
	number_two := 2
	
	sum_numbers := add_numbers(number_one, number_two)


	fmt.Println(sum_numbers)

}
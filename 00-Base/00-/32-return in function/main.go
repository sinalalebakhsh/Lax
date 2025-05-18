package main

import "fmt"

func return_sum(number_one, number_two int) int {
	sum_numbers := number_one + number_two
	return sum_numbers
}

func main() {
	fmt.Println(return_sum(2, 5))
}
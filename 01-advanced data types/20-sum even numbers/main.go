package main

import (
	"fmt"

)

func get_number_new(number int) []int {
	var list_even_numbers []int

	for {
		if number == 0 {
			break
		} else {
			if number % 2 == 0 {
				list_even_numbers = append(list_even_numbers, number)
				number--
			} else if number % 2 != 0 {
				number--
			}
		}
	}

	return list_even_numbers
}

func print_sum_all_even_numbers(numbers ...int) int {
	var last_number int
	for _, v := range numbers {
		if v%2 == 0 {
			last_number += v
		}
	}
	return last_number
}


func main() {

	get_all_even_numbers := get_number_new(22)
	fmt.Println(get_number_new(22))

	fmt.Println(print_sum_all_even_numbers(get_all_even_numbers...))
	
}

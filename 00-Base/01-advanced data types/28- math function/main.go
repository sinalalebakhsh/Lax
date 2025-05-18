package main

import "fmt"

type mathFunc func(int) int 

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}


func apply_function(number int, function mathFunc)  {
	result := function(number)
	fmt.Println(result)
}

func main()  {
	apply_function(2, double)
	apply_function(2, triple)

}
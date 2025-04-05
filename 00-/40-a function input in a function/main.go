package main

import "fmt"

func for_return_two_strings(name string) string  {
	return "Hello " + name
}

func for_print(a_function string) {
	fmt.Println(a_function)
}

func main() {
	for_print(for_return_two_strings("Sina"))
}
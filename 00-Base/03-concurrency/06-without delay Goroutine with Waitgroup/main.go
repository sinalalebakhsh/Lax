package main

import "fmt"

func main() {
	var a, b, c int

	go func() {
		a = 1
	}()

	go func() {
		b = 2
	}()

	go func() {
		c = 3
	}()

	fmt.Println(a, b, c)
}

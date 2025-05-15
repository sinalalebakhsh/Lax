package main

import (
	"fmt"
	"time"
)



func main()  {
	print_letters()
	fmt.Println()
	print_numbers()
	fmt.Println()
}

func print_letters()  {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), " ")
		time.Sleep(time.Millisecond * 50)
	}
}

func print_numbers()  {
	for i := 1; i <= 50; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Millisecond * 50)
	}
}
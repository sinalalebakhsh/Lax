package main

import (
	"fmt"
	"time"
)



func main()  {
	
	go print_letters()
	fmt.Println()
	go print_numbers()
	fmt.Println()
	// time.Sleep(time.Millisecond * 5000)
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
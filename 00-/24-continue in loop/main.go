package main

import "fmt"

func main()  {
	for i := 0; i < 4; i++ {
		fmt.Print("B ")
		continue
		fmt.Print("S ")
	}
	fmt.Println()
}
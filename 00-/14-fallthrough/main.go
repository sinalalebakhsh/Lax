package main

import "fmt"

func main()  {
	var A_number int16 = 3
	switch A_number {
	case 1:
		fmt.Println(1)
		fallthrough
	case 3:
		fmt.Println(3)
		fallthrough
	case 4:
		fmt.Println(4)
		// fallthrough
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	case 7:
		fmt.Println(7)
	}
	
}
package main

import "fmt"

func main()  {
	ages := map[string]int{
		"sina": 1,
		"lina": 2,
		"wina": 3,
		"qina": 4,
		"zina": 5,
	}
	
	fmt.Println("before clear")
	fmt.Println(ages)

	clear(ages)

	fmt.Println("after clear")
	fmt.Println(ages)
	
}
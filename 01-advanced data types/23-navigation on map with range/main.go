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
	
	for keys, values := range ages {
		fmt.Println(keys, ":", values)
	}
}
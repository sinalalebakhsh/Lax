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
	
	item, ok := ages["LAX"]
	if ok {
		fmt.Println(item)
	} else {
		fmt.Println("LAX is not in ages")
		fmt.Println(ages)
	}

	
}
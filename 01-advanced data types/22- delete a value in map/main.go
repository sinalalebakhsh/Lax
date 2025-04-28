package main

import "fmt"

func main()  {

	another_map_value := map[string]int{
		"sina": 100,
		"lina": 150,
	}
	fmt.Println(another_map_value)

	delete(another_map_value, "sina")
	fmt.Println(another_map_value)

	
}
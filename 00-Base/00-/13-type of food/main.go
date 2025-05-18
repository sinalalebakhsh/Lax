package main

import "fmt"

func main()  {
	var food string = "French Fries"

	switch food {
	case "Coke", "Water", "Tea", "Coffee":
		fmt.Println("You ordered a drink.")		
	case "Pizza", "Burger", "Sandwich":
		fmt.Println("You ordered a main food.")		
	case "Salad", "French Fries", "Yogurt":
		fmt.Println("You ordered a side dish.")		
	default:
		fmt.Println("Unknown food.")		
	}


}
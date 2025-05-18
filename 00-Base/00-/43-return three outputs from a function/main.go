package main

import "fmt"

func pizza() (string, int, string) {
	return "Pizza", 350000, "Cheese + Peperoni"
}

func burger() (string, int, string) {
	return "Burger", 185000, "Meat + Tomato + Cheese"
}

func kebab() (string, int, string) {
	return "Kebab", 250000, "Meat + Rice"
}

func show_detail(f func() (string, int, string)) {
	name, price, ingredients := f()
	fmt.Println("Show Food Details:")
	fmt.Println("Name:", name)
	fmt.Println("Price:", price)
	fmt.Println("Ingredients:", ingredients)
	fmt.Println()
}

func main() {
	show_detail(pizza)
	show_detail(burger)
	show_detail(kebab)
}



package main

import "fmt"

func main()  {
	var food string = "pizza"
	var amount = 500000

	switch {
	case amount >= 500000:
		fmt.Println("30% discount")
		amount = int(float64(amount) * 0.7)
	case food == "burger":
		fmt.Println("20% discount")
		amount = int(float64(amount) * 0.8)
	case food == "pizza":
		fmt.Println("10% discount")
		amount = int(float64(amount) * 0.9)

	}
	fmt.Println(amount, "Toman")
}
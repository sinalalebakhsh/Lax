package main

import "fmt"

func my_print(name string) {
	name = "sina"
	fmt.Println(name)
}

func main()  {
	var name string = "GGG"

	fmt.Println(name)

	my_print(name)

	fmt.Println(name)

}

/* Output:


	GGG
	sina
	GGG


*/



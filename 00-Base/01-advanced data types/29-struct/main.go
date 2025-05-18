package main

import "fmt"

type user struct {
	Name string
	Age int
	Email string
}

func main()  {

	user_1 := user{}

	user_1.Name = "sina"
	user_1.Age = 33
	// user_1.Email = "sina@gmail.com"

	fmt.Println(user_1.Name)
	fmt.Println(user_1.Age)
	fmt.Println("-",user_1.Email, "-")
	
}
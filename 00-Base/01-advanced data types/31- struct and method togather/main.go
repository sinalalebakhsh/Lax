package main

import "fmt"

type user struct {
	first_name string
	last_name  string
	age        int
}

func (u user) full_name() string  {
	return u.first_name + " " + u.last_name
}

func (u user) print()  {
	fmt.Println("Name:", u.full_name())
	fmt.Println("Age:", u.age)
}

func main() {
	sina := user{
		first_name: "sina",
		last_name: "lalehbakhsh",
		age: 33,
	}

	sina.print()
	
}

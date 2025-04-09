package main

import "fmt"

func main() {

	student_names := [20]string{}

	var student_scores [20]int

	// default values is " " a empty string
	fmt.Println(student_names)

	// default value is 0
	fmt.Println(student_scores)

	// define with value in same time
	student_lastnames := [20]string{
		"lalebakhsh",
		"loolebakhsh",
		"jalebakhsh",
		"joolebakhsh",
		"balebakhsh",
		"boolebakhsh",
	}

	// define six first lastnames 
	fmt.Println(student_lastnames)
}

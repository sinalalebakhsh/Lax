package main

import "fmt"


// Can't write this:
// Variable_name := "aalkjsdnz"
// You can not used this variable ******
// Can do this:
var Variable_name = "aalkjsdnz"


func main() {
	var اسم string
	var age int
	var gender string
	// This is Hello World Application
	اسم = "سینا"
	age = 32
	gender = "مرد"

	
	fmt.Println("سلام", اسم)
	fmt.Println("سن شما" , age , "سال هست.")
	fmt.Println("و جنسیت شما", gender, " می باشد")

	myLove := "bahareh"
	fmt.Println(myLove)
}


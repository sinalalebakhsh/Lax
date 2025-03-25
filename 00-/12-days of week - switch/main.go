package main

import "fmt"


func main()  {
	var days int8 = 0

	switch days {
	case 0:
		fmt.Println("today is Shanbe")
	case 1:
		fmt.Println("today is Yekshanbe")
	case 2:
		fmt.Println("today is Doshanbe")
	case 3:
		fmt.Println("today is Seshanbe")
	case 4:
		fmt.Println("today is Chaharshanbe")
	case 5:
		fmt.Println("today is Panjshanbe")
	case 6:
		fmt.Println("today is Jomee")
	default:
		fmt.Println("Nadarim hamchin roozi !!!")
	}
	
}
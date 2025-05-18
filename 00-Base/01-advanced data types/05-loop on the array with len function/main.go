package main

import "fmt"

func main()  {

	student_scores := [5]int{0,1,2,3,4}

	for i := 0; i < len(student_scores); i++ {
		fmt.Println("student score:",student_scores[i])
	}
	
}
package main

import "fmt"

func main()  {
	student_scores := make([]int, 8, 16)

	fmt.Println("length:",len(student_scores))
	fmt.Println("capacity:",cap(student_scores))
	fmt.Println(student_scores)	
	fmt.Println("======================")	
	student_scores = append(student_scores, 2)
	fmt.Println("length:",len(student_scores))
	fmt.Println("capacity:",cap(student_scores))
	fmt.Println(student_scores)	


	fmt.Println("======================")	
	student_scores = append(student_scores, 88,3,4,5,6,7,8,9,10,11,12,13,14,15,16)
	fmt.Println("length:",len(student_scores))
	fmt.Println("capacity:",cap(student_scores))
	fmt.Println(student_scores)	

}
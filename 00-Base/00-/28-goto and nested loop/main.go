package main

import "fmt"

func main()  {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(j)
			if j == 3 {
				goto finish
			}
		}
	}
	finish:
	fmt.Println("finish is here")
}
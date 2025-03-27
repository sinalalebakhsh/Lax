package main

import "fmt"

// مضرب اعداد
// چاپ میشود

func main()  {
	var i int64
	var j int64
	for i = 2; i <= 10; i++ {
		fmt.Println("multiples of number:", i)
		for j = i; j <= 20; j++ {
			if j%i==0 {
				fmt.Print(j, " ")
			}
		}
		fmt.Println("\n-------------------")
	}
	
}
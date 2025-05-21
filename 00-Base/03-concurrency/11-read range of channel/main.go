package main

import "fmt"

func main()  {

	my_channel := make(chan string, 2)
	my_channel <- "sina"
	my_channel <- "lalex"
	// for per_words := range my_channel {
	// 	fmt.Println(per_words)
	// }
	// OR
	for i := 0; i < 2; i++ {
		fmt.Println(<-my_channel)
	}
	
}



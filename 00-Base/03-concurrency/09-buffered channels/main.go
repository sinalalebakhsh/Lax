package main

import (
	"fmt"
)

func main() {
	my_channel := make(chan string, 2)

	go func() {
		my_channel <- "string-22"
	}()
	
	fmt.Println(<-my_channel)
	
	my_channel <- "string-1"

	fmt.Println(<-my_channel)

}

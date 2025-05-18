package main

import (
	"fmt"
	"time"
)

func hello_world(first_channel chan string) {
	time.Sleep(5 * time.Second)
	first_channel <- "hello world"
}

func main() {
	my_channel := make(chan string)

	go hello_world(my_channel)

	fmt.Print("waiting ")
	time.Sleep(time.Second)
	fmt.Print(".")
	time.Sleep(time.Second)
	fmt.Print(".")
	time.Sleep(time.Second)
	fmt.Print(".")
	time.Sleep(time.Second)
	fmt.Println()
	message := <-my_channel

	fmt.Println("got channel:", message)

}

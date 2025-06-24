package main

import (
	"fmt"
	"time"
)

func main() {
	coffee_chanel := make(chan string)
	tea_chanel := make(chan string)
	food_chanel := make(chan string)
	quit := make(chan bool)
	
	go func()  {
		time.Sleep( time.Second)
		tea_chanel <- "Black Tea"
	}()
	

	go func()  {
		time.Sleep(2 *time.Second)
		coffee_chanel <- "Esspresso coffee"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		quit <- true
	}()

	defer fmt.Println("Berid khoonatoon Shop is closed!")

	for {
		select {
		case coffee := <-coffee_chanel:
			fmt.Printf("Received a coffee order: %s\n", coffee)
		case tea := <-tea_chanel:
			fmt.Printf("Received a tea order: %s\n", tea)
		case food := <-food_chanel:
			fmt.Printf("Received a food order: %s\n", food)
		case <-quit:
			return
		default:
			fmt.Print()
		}
	}


}

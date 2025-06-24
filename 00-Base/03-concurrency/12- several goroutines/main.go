package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	coffee_chanel := make(chan string)
	tea_chanel := make(chan string)
	food_chanel := make(chan string)

	go func()  {
		time.Sleep(time.Second)
		tea_chanel <- "Black Tea"
	}()

	var wait_group sync.WaitGroup
	wait_group.Add(1)

	go func() {
		coffee := <-coffee_chanel
		fmt.Printf("Received a coffee order: %s\n", coffee)
		wait_group.Done()
	}()
	go func() {
		tea := <-tea_chanel
		fmt.Printf("Received a tea order: %s\n", tea)
		wait_group.Done()
	}()
	go func() {
		food := <-food_chanel
		fmt.Printf("Received a food order: %s\n", food)
		wait_group.Done()
	}()
	
	wait_group.Wait()

}

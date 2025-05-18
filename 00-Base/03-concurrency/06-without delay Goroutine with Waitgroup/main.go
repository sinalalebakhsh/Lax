package main

import (
	"fmt"
	"sync"
)

func main() {
	var a, b, c int

	var wait_group sync.WaitGroup

	wait_group.Add(3)



	go func() {
		a = 1
		wait_group.Done()
	}()

	go func() {
		b = 2
		wait_group.Done()
	}()

	go func() {
		c = 3
		wait_group.Done()
	}()
	
	wait_group.Wait()

	fmt.Println(a, b, c)
}

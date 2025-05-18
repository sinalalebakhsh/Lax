package main

import (
	"fmt"
	"sync"
	"time"
)

func print_letters()  {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), " ")
		time.Sleep(time.Millisecond * 50)
	}
}

func print_numbers()  {
	for i := 1; i <= 50; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Millisecond * 50)
	}
}


func main(){
	var wait_group sync.WaitGroup
	wait_group.Add(2)

	go func()  {
		print_letters()
		wait_group.Done()
	}()

	go func()  {
		print_numbers()	
		// wait_group.Done()
	}()

	wait_group.Wait()
	
	fmt.Println("finished")
		
}

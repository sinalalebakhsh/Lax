package main

import (
	"fmt"
	"time"
)

func heavy_task(my_channel chan int)  {
	time.Sleep(time.Second)
	i := <- my_channel
	fmt.Println("task:", i, "done.")
}


func main()  {
	
	a_channel := make(chan int, 3)
	for i := 0; i <= 100; i++ {
		a_channel <- i
		go heavy_task(a_channel)
	}


}


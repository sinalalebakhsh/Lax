package main

import (
	"fmt"
	"sync"
	"time"
)


func main()  {
	const number_loop = 15

	var wait_groups sync.WaitGroup

	wait_groups.Add(number_loop)
	
	t := time.Now()

	for i := 0; i < number_loop; i++ {
		go func(i int)  {
			to_do_task(i, time.Millisecond*time.Duration(i))
			wait_groups.Done()
		}(i)
	}

	wait_groups.Wait()
	
	fmt.Println("total:", time.Since(t))

}

func to_do_task(i int, d time.Duration)  {
	time.Sleep(d)
	fmt.Println("task", i, "finished")
}

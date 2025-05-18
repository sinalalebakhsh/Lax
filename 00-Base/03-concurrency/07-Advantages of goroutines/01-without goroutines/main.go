package main

import (
	"fmt"
	"time"
)


func main()  {
	t := time.Now()
	const number_loop = 15

	for i := 0; i < number_loop; i++ {
		to_do_task(i, time.Millisecond*time.Duration(i))
	}

	fmt.Println("total:", time.Since(t))

}

func to_do_task(i int, d time.Duration)  {
	time.Sleep(d)
	fmt.Println("task", i, "finished")
}

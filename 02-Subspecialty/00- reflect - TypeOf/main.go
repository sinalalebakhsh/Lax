package main

import (
	"fmt"
	"reflect"
)

func main()  {
	numbers := []int{1,2,3,4}
	copy_to_numbers := copy(numbers, []int{77, 99})

	fmt.Println(reflect.TypeOf(copy_to_numbers))	

}
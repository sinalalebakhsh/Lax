package main

import "fmt"

var is_global_scope = 5

func main() {

	var is_global_scope = 10

	fmt.Println(is_global_scope)
}
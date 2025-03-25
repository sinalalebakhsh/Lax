// عملگرهای بیتی

package main

import "fmt"

func main() {
	var a int = 0
	var b int = 1

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	fmt.Println("a & b = ", a&b)
	fmt.Println("a | b = ", a|b)
	fmt.Println("a ^ b = ", a^b)
	fmt.Println("^a = ", ^a)
	fmt.Println("^b = ", ^b)
	fmt.Println("a >> 2 = ", a>>2)
	fmt.Println("a << 2 = ", a<<2)
	fmt.Println("b >> 2 = ", b>>2)
	fmt.Println("b << 2 = ", b<<2)
}

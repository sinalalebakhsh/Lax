package main

import "fmt"

func main() {
	var a int = 2
	var b int = 10

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	fmt.Println("(a+b) ", a, "+", b, "=",   a+b)
	fmt.Println("(b-1) ", b, "-", "1", "=", b-1)
	fmt.Println("(a*b) ", a, "*", b, "=",   a*b)
	fmt.Println("(b/a) ", b, "/", a, "=",   b/a)
	fmt.Println("(b%a) ", b, "%", a, "=",   b%a)

	b++
	fmt.Println("b++ = ",b)

	a--
	fmt.Println("a-- =",a)

}

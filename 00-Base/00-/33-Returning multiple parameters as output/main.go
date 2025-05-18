package main

import "fmt"

const pi = 3.14

func AreaAndPerimeter(r float64) (float64, float64) {
	perimeter := pi * (r * 2)
	area := pi * (r * r)
	return perimeter, area
}

func main() {

	area, perimeter := AreaAndPerimeter(1)
	fmt.Println("Area of circle is:", area)
	fmt.Println("Perimeter of circle is:", perimeter)

}
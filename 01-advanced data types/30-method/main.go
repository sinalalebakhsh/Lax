package main

import "fmt"

type dayere float64

func (d dayere) mohit() float64 {
	shoaA := float64(d)
	return shoaA * shoaA * 3.14
	
}

func (d dayere) masahat() float64  {
	shoA := float64(d)
	return (shoA * 2) * 3.14
}

func main()  {
	var dayereye_man dayere = 2 	

	mohitesh := dayereye_man.mohit()
	masahatesh := dayereye_man.masahat()

	fmt.Println("mohitesh:", mohitesh)
	fmt.Println("masahatesh:",masahatesh)
}
package main

import "fmt"

type player struct {
	name string
	hp   int
}

func (p *player) hit() {
	p.hp--
}

func main() {
	player1 := player{
		name: "sina",
		hp: 100,
	}
	fmt.Println("before hit:", player1.hp)
	player1.hit()
	fmt.Println("after hit:", player1.hp)
	player1.hit()
	fmt.Println("after hit:", player1.hp)
	player1.hit()

}

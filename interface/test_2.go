package main

import (
	"fmt"
)

type Resetter interface {
	Reset()
}

type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{0, 0}
}

func Reset(r Resetter) {
	r.Reset()
}

func main() {

	player := Player{50, Coordinate{5, 5}}
	Reset(&player)
	fmt.Println(player)
}

// +build OMIT

package main

import "fmt"

type player struct {
	nickname string
	health   int
}

func (p player) warcry() string {
	return "I don't want to die :("
}

func (p player) consumePotion() {
	p.health = 100
}

func main() {
	p := player{
		nickname: "Urgoth the Destroyer",
		health:   12,
	}

	p.consumePotion()

	if p.health == 100 {
		fmt.Println(p.nickname, "lives to fight another day")
	} else {
		fmt.Println(p.nickname, "still suffers")
	}
}

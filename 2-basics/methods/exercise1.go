// +build OMIT

package main

import "fmt"

type player struct {
	nickname string
}

func (p player) warcry() string {
	return "I don't want to die :("
}

func main() {
	p := player{
		nickname: "Urgoth the Destroyer",
	}

	fmt.Println("Fearlessly", a.nickname, "dived into the abyss, screaming:", p.warcry())
}

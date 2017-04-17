// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
type sentence string

func (s sentence) Scream() {
	fmt.Println(s + "!")
}

func (s sentence) Ask() {
	fmt.Println(s + "?")
}

func main() {
	var s sentence
	s = "Go"

	s.Ask()
	s.Scream()
}

// END OMIT

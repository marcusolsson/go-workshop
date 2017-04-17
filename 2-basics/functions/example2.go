// +build OMIT 

package main

import "fmt"

func main() {
	var fn func(i int) bool

	fn = func(i int) bool {
		return i < 0
	}

	b1 := fn(-4)
	b2 := fn(11)

	fmt.Println(b1)
	fmt.Println(b2)
}

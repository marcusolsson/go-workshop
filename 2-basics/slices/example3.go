// +build OMIT

package main

import "fmt"

func main() { // OMIT
	letters := []string{
		"A", "B", "C",
	}

	for i := 0; i < len(letters); i++ {
		fmt.Println(i, letters[i])
	}

	for idx, letter := range letters {
		fmt.Println(idx, letter)
	}
} // OMIT

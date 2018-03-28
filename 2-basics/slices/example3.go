// +build OMIT

package main

import "fmt"

func main() {
	letters := []string{
		"A", "B", "C",
	}

	for idx := 0; idx < len(letters); idx++ {
		fmt.Println(idx, letters[idx])
	}

	for idx, letter := range letters {
		fmt.Println(idx, letter)
	}
}

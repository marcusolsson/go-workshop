// +build OMIT

package main

import "fmt"

func main() {
	// START EX1 OMIT
	arr := [5]int{1, 2, 3, 4, 5}

	var s []int = arr[1:4]

	fmt.Println(s)
	// END EX1 OMIT

	// START EX2 OMIT
	elems := []int{1, 2, 3} // [1 2 3]
	// END EX2 OMIT

	// START EX3 OMIT
	elems := make([]int, 3) // [0 0 0]
	// END EX3 OMIT
}

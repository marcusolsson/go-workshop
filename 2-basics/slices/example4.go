// +build OMIT

package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	var s []int = arr[1:4]

	length := len(s) // HL

	fmt.Println(length)
}

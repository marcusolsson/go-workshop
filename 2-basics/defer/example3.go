// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("entering main")
	
	defer func() {
		fmt.Println("exiting main")
	}()

	fmt.Println("doing work")
}

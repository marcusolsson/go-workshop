package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

func example() {
	// START DECL OMIT
	var ptr *user // nil
	// END DECL OMIT

	// START REF OMIT
	var val user
	ptr = &val // HL

	fmt.Printf("val = %v; ptr = %v", val, ptr)
	// END REF OMIT

	// START DEREF OMIT
	var newval user
	newval = *ptr // HL
	// END DEREF OMIT
}

package main

type user struct {
	firstName string
	lastName  string
}

func main() {
	// START DECL OMIT
	var ptr *user // nil
	// END DECL OMIT

	// START REF OMIT
	var val user
	ptr = &val // HL
	// END REF OMIT

	// START DEREF OMIT
	var newval user
	newval = *ptr // HL
	// END DEREF OMIT

	_ = newval
}

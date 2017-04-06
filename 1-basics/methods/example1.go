package main

import "fmt"

// START METHOD OMIT
type user struct {
	firstName string
	lastName  string
}

func (u user) fullName() string { // HL
	return u.firstName + " " + u.lastName
}

// END METHOD OMIT

func main() {
	u := user{
		firstName: "Jane",
		lastName:  "Doe",
	}

	fmt.Println(u.fullName()) // HL
}

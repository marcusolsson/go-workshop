// +build OMIT

package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

func (u user) fullName() string { // HL
	return u.firstName + " " + u.lastName
}

func (u *user) rename(first, last string) { // HL
	u.firstName = first
	u.lastName = last
}

func main() {
	u := user{
		firstName: "Jane",
		lastName:  "Doe",
	}
	u.rename("Bob", "Doe")

	fmt.Println(u.fullName())
}

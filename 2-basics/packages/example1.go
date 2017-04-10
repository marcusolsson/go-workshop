// START HELLO OMIT
package main // HLpackage

import "fmt" // HLimport

func main() {
	fmt.Println("Hello, world!") // HLuse
}

// END HELLO OMIT

// START EXPORT OMIT
func Hello(n string) string {
	return unexportedHelper(n)
}

func unexportedHelper(n string) string {
	return "Hello, " + n
}

// END EXPORT OMIT

// START STRUCT OMIT

type User struct {
	FirstName string
	LastName  string
	modified  bool
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// END STRUCT OMIT

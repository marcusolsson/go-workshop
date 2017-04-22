package error

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

// START MYERROR OMIT
type myError struct {
	reason string
}

func (e myError) Error() string {
	return e.reason
}

// END MYERROR OMIT

func add(x, y int) (int, error) {
	if x > 10 && y > 10 {
		return 0, errors.New("number too big")
	}
	return x + y, nil
}

func example() {
	// START HANDLE OMIT
	sum, err := add(4, 5)
	if err != nil {
		fmt.Println("something bad happened")
	}
	// END HANDLE OMIT
}

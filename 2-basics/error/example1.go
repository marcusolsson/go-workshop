// +build OMIT

package error

type error interface {
	Error() string
}

type myError struct {
	reason string
}

func (e myError) Error() string {
	return e.reason
}

func add(x,y) (int, error) {
	if x > 10 && y > 10 {
		return 0, errors.New("number too big")
	}
	return x + y
}
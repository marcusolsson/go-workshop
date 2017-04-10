// +build OMIT

package main

func hello(name string) string {
	return "Hello, " + name + "!"
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

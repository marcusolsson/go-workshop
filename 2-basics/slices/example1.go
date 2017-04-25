package main

func example1() {
	// START EX1 OMIT
	var arr [5]int // [0 0 0 0 0]

	arr[4] = 100 // [0 0 0 0 100]

	i := arr[4] // 100
	// END EX1 OMIT

	// START EX2 OMIT
	arr = [5]int{1, 2, 3, 4, 5}
	// END EX2 OMIT

	// START EX3 OMIT
	n := len(arr)
	// END EX3 OMIT
}

package main

func example1() { // OMIT
	var arr [5]int // [0 0 0 0 0]

	arr[4] = 100 // [0 0 0 0 100]
} // OMIT

func example2() { // OMIT
	arr := [5]int{1, 2, 3, 4, 5}
} // OMIT

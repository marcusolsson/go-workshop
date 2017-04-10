package main

func main() {
	// START DECL OMIT
	var name string
	var year int
	var ratio float64
	var valid bool
	// END DECL OMIT

	// START ZERO OMIT
	var name string   // ""
	var year int      // 0
	var ratio float64 // 0.0
	var valid bool    // false
	// END ZERO OMIT

	// START INIT OMIT
	var name string = "Beth"
	var year int = 2008
	var ratio float64 = 1.3
	var valid bool = true
	// END INIT OMIT

	// START INFER OMIT
	var name = "Beth"
	var year = 2008
	var ratio = 1.3
	var valid = true
	// END INFER OMIT

	// START SHORT OMIT
	name := "Beth"
	year := 2008
	ratio := 1.3
	valid := true
	// END SHORT OMIT
}

package main

type movie struct {
	title string
	year  int
}

func example1() {
	// START INIT OMIT
	m := movie{
		title: "A New Hope",
		year:  1999,
	}
	// END INIT OMIT

	// START GETSET OMIT
	m.title = "Star Wars IV: A New Hope"
	t := m.title
	// END GETSET OMIT
}

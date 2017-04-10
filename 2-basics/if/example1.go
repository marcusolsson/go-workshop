// +build OMIT

package main

func status(points int) string {
	// START IF OMIT
	if points > 50 {
		return "passed"
	}
	// END IF OMIT
	return "failed"
}

func review(points int) string { 
	if points > 60 {
		return "good"
	} else if points > 30 {
		return "neutral"
	} else {
		return "bad"
	}
}

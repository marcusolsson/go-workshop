package main

import "fmt"

type notifier interface {
	notify()
}

var n notifier // nil

// START IMPL OMIT
type emailNotifier struct {
	email string
}

func (n emailNotifier) notify() {
	fmt.Println("sending email to", n.email)
}

// END IMPL OMIT

// START EXAMPLE OMIT
func main() {
	en := emailNotifier{
		email: "jane@doe.com",
	}

	sendNotification(en)
}

func sendNotification(n notifier) { // HL
	n.notify()
}

// END EXAMPLE OMIT

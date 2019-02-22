package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("I'm goroutine")
	// Sends and receives to a channel are blocking by default
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	// Blocking here, which means that until some Goroutine writes data to the `done`
	// channel, the control will not move to the next line of code. This line
	// receives data from the done channel but does not use or store that data
	// in any variable
	<-done
	fmt.Println("I'm main")
}

package main

import (
	"log"

	deque "github.com/bsnux/go-cookbook/DequeGeneric"
)

// Persona represents one person
type Persona struct {
	Age  int
	Name string
}

func main() {
	d := deque.NewDeque()
	d.PushRight(&Persona{1, "Bob"})
	d.PushRight(&Persona{2, "Alice"})
	d.PushLeft(&Persona{3, "Charlie"})
	log.Println(d)
	log.Println(d.PopRight())
	log.Println(d)
	log.Println(d.PopLeft())
	log.Println(d)
}

package main

import (
	"log"

	"github.com/bsnux/go-cookbook/deque"
)

func main() {
	d := deque.NewDeque()
	d.PushRight(1)
	d.PushRight(2)
	d.PushLeft(3)
	log.Println(d)
	log.Println(d.PopRight())
	log.Println(d)
	log.Println(d.PopLeft())
	log.Println(d)
}

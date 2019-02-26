// `defer` is especially useful for releasing resources that must be released
// under any conditions. It is useful for many reasons, the most common of which
// are to close an open connection or unlock a Mutex immediately before the function ends
package main

import (
	"log"
)

func main() {
	log.Println("first")
	log.Println("second")
	defer log.Println("e")

	log.Println("a")
	log.Println("b")
	log.Println("c")

	printThis("d")

	// When a function returns, its deferred calls are executed in LIFO order.
	for i := 0; i < 10; i++ {
		defer log.Println(i)
	}
}

func printThis(str string) {
	log.Println(str)
}

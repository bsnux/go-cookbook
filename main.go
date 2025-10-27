package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// `counter` is global shared state that needs to be protected using a mutex
	mutex.Lock()
	defer mutex.Unlock()

	counter++
}

func main() {
	var wg sync.WaitGroup
	var top = 1000

	// Creating goroutines to increment the counter
	for i := 0; i < top; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

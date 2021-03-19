/*
* Simple example for doing concurrent request without using channels.
* Channels are useful for sharing information between goroutines, but
* we don't need to do that in this example
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// User represents info from a single user
type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/4",
	}

	start := time.Now()
	wg.Add(len(urls))
	for _, url := range urls {
		go makeRequest(url)
	}

	wg.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

// makeRequest executes a HTTP request
func makeRequest(url string) error {
	defer wg.Done()

	resp, _ := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var user User
	json.Unmarshal(body, &user)
	fmt.Printf("ID: %d \tTitle: %s \tCompleted: %t\n", user.ID, user.Title, user.Completed)

	return nil
}

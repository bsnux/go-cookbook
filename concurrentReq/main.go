package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// User represents info from a single user
type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/4",
	}

	start := time.Now()
	ch := make(chan bool)
	for _, url := range urls {
		go makeRequest(url, ch)
	}

	for range urls {
		<-ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

// makeRequest executes a HTTP request
func makeRequest(url string, ch chan<- bool) {
	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)
	var user User
	json.Unmarshal(body, &user)
	fmt.Printf("ID: %d \tTitle: %s \tCompleted: %t\n", user.ID, user.Title, user.Completed)
	ch <- true
}

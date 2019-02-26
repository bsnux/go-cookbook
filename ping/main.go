// Simple web server listening on port 8080
//
// Launch this program, open a new terminal, then type:
// `curl -L http://localhost:8080`
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
	log.Printf("User-agent: %s\n", r.UserAgent())
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

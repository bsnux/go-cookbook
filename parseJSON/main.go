// How to parse JSON when we don't know about what data are we going to receive
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	var result interface{}

	url := "https://jsonplaceholder.typicode.com/todos/"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Simple way to parse. We know we're getting an array
	var results []map[string]interface{}
	json.Unmarshal([]byte(body), &results)
	for _, record := range results {
		fmt.Printf("Title: %s\n", record["title"])
	}

	fmt.Println("**********************")

	// Another way to parse
	switch result.(type) {
	case []interface{}:
		// We're expecting this type but `switch` with different types is the way
		// to deal when we don't know about the JSON result
		arr := result.([]interface{})
		for _, item := range arr {
			record := item.(map[string]interface{})
			fmt.Printf("Title: %s\n", record["title"])
		}
		fmt.Printf("Total records: %d\n", len(arr))
	default:
		log.Fatal("JSON type not understood")
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

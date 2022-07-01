package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/bsnux", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error sending http request: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var result map[string]interface{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		fmt.Printf("json: could not parse json: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result["bio"])
}

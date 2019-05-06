package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Printing out request time for a given URL\n\nUsage %s [options] <url>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	var intPtr = flag.Int("c", 1, "number of concurrent request")

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	ch := make(chan string)
	chSecs := make(chan float64)
	for i := 0; i < *intPtr; i++ {
		go doRequest(flag.Args()[0], ch, chSecs)
	}

	var secs float64

	for i := 0; i < *intPtr; i++ {
		log.Printf(<-ch)
		secs += <-chSecs
	}

	log.Printf("Average: %.2f seconds", secs/float64(*intPtr))
}

func doRequest(url string, ch chan<- string, chSecs chan<- float64) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	ch <- fmt.Sprintf("Request done in %.2f seconds. HTTP: %d", elapsed.Seconds(), res.StatusCode)
	chSecs <- elapsed.Seconds()
}

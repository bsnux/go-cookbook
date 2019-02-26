package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	dir := "/tmp/"
	var files int

	// Walking through directories using an anonymous function
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}
		if !info.IsDir() {
			files++
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	log.Printf("Total files: %d", files)
	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

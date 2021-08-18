package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "10")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Printf("Done!")
}

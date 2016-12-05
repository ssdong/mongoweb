package main

import (
	"log"
	"os/exec"
)

func main() {
	// Check if there is already a mongodb instance running
	cmd := exec.Command("mongo")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

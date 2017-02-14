package main

import (
	"flag"
	"fmt"
	"log"
	"mongoweb/server"
	"os"
	"os/exec"
)

var usage = `
  Usage: mongoweb [options]

  -port <port>     The port that mongoweb listens to
`

func main() {
	if len(os.Args) < 2 {
		man()
	}

	var port string

	flag.StringVar(&port, "port", "", "The port that mongoweb listens to")

	flag.Parse()

	// Check if there is only a MongoDB instance running
	cmd := exec.Command("mongo")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error: MongoDB needs to run before this client")
	}

	// @TODO Check if port is available
	server.Listen(":" + port)
}

// This will print out the flag options
func man() {
	fmt.Printf("%s\n", usage)
	os.Exit(0)
}

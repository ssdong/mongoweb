package main

import (
	"log"
	"mongoweb/command"
	"os/exec"
)

func main() {
	//Ensure that MongoDB is running before attempting to start the mongo shell
	cmd := exec.Command("mongo")
	err := cmd.Run()
	if err != nil {
		log.Fatal("MongoDB needs to run before mongo shell")
	}

	commandManager := new(command.CmdManager)
	commandManager.GetDbs()
}

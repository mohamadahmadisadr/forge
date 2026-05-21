package main

import (
	"forge/internal/cli"
	"forge/internal/task"
	"log"
	"os"
)

func main() {

	args, err := cli.ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
	data, err := task.GetYamlConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = cli.RunCommand(args, data)
	if err != nil {
		log.Fatal(err)
	}

}

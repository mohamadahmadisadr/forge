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
	}

	conf, err := task.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.RunCommand(args, conf); err != nil {
		log.Fatal(err)
	}
}

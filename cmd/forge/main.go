package main

import (
	"forge/internal/cli"
	"log"
	"os"
)

func main() {
	args, err := cli.ParseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cli.RunCommand(args)
	if err != nil {
		log.Fatal(err)
	}

}

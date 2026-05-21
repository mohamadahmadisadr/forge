package cli

import (
	"fmt"
)

type Cli struct {
	Command string
	Task    string
	Args    []string
}

func ParseArgs(args []string) (*Cli, error) {

	if len(args) < 2 {
		return nil, fmt.Errorf("cmd not found")
	}

	if len(args) < 3 {
		return nil, fmt.Errorf("task not found")
	}
	var arguments []string

	if len(args) > 3 {
		arguments = args[3:]
	}

	return &Cli{
		Command: args[1],
		Task:    args[2],
		Args:    arguments,
	}, nil
}

package cli

import "fmt"

type CLI struct {
	Command string
	Target  string
}

func ParseArgs(args []string) (*CLI, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("command required")
	}

	if len(args) < 3 {
		return nil, fmt.Errorf("target required")
	}

	return &CLI{
		Command: args[1],
		Target:  args[2],
	}, nil
}

package cli

import (
	"fmt"
	"forge/internal/executor"
)

func RunCommand(cli *Cli) error {
	switch cli.Command {
	case "run":
		fmt.Printf("running: %s", cli.Task)
		executor.Execute(cli.Task, cli.Args)

	default:
		return error(fmt.Errorf("Unknown command: %s", cli.Command))
	}
	return nil
}

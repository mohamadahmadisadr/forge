package cli

import (
	"fmt"
	"forge/internal/executor"
	"forge/internal/task"
)

func RunCommand(cli *Cli, conf *task.YamlConfig) error {
	switch cli.Command {
	case "run":
		fmt.Printf("running: %s", cli.Task)
		err := executor.Execute(conf.Tasks[cli.Task])
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown command: %s", cli.Command)
	}
	return nil
}

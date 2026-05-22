package cli

import (
	"fmt"
	"forge/internal/executor"
	"forge/internal/task"
	"sync"
)

func RunCommand(cli *Cli, conf *task.YamlConfig) error {
	switch cli.Command {
	case "run":
		fmt.Printf("running: %s\n", cli.Task)
		if cli.Task == "all" {
			return executeWorkflow(conf)
		}
		mapData, ok := conf.Tasks[cli.Task]
		if !ok {
			return fmt.Errorf("mapData not found: %s", cli.Task)
		}

		err := executeTask(mapData)
		if err != nil {
			return fmt.Errorf("failed running %s: %w", mapData.Command, err)
		}

	default:
		return fmt.Errorf("unknown command: %s", cli.Command)
	}
	return nil
}

func executeWorkflow(conf *task.YamlConfig) error {
	var wg sync.WaitGroup
	all, ok := conf.Tasks["all"]
	if !ok {
		return fmt.Errorf("task 'all' not defined")
	}
	for _, name := range all.Parallel {
		data, ok := conf.Tasks[name]
		if !ok {
			return fmt.Errorf("sub task does not exist: %s", data.Command)
		}
		wg.Add(1)
		go func(t task.Task) {

			defer wg.Done()
			if err := executeTask(t); err != nil {
				fmt.Printf("failed: %s -> %s\n", t.Command, err)
			}

		}(data)

	}
	wg.Wait()
	return nil
}

func executeTask(task task.Task) error {
	err := executor.Execute(task)
	if err != nil {
		return err
	}
	return nil
}

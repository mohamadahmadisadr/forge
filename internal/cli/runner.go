package cli

import (
	"fmt"
	"forge/internal/executor"
	"forge/internal/planner"
	"forge/internal/task"
)

func RunCommand(cli *CLI, conf *task.Config) error {
	switch cli.Command {
	case "run":
		return runWorkflow(cli.Target, conf)

	default:
		return fmt.Errorf("unknown command: %s", cli.Command)
	}
}

func runWorkflow(target string, conf *task.Config) error {
	graph := planner.BuildGraph(conf)

	executionPlan, err := planner.ResolveExecutionPlan(graph, target)
	if err != nil {
		return err
	}

	fmt.Println("execution plan:")

	for _, node := range executionPlan {
		fmt.Printf(" - %s\n", node.Name)
	}

	for _, node := range executionPlan {
		fmt.Printf("running: %s\n", node.Name)

		if err := executor.Execute(node); err != nil {
			return fmt.Errorf("task failed [%s]: %w", node.Name, err)
		}
	}

	return nil
}

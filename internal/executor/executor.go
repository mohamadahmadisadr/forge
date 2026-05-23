package executor

import (
	"bufio"
	"fmt"
	"forge/internal/task"
	"os/exec"
)

func Execute(node *task.TaskNode) error {
	cmd := exec.Command(node.Command, node.Args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go streamOutput(stdout, node.Name)
	go streamOutput(stderr, node.Name)

	return cmd.Wait()
}

func streamOutput(pipe interface {
	Read([]byte) (int, error)
}, taskName string) {

	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", taskName, scanner.Text())
	}
}

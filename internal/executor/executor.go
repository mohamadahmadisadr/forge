package executor

import (
	"bufio"
	"fmt"
	"forge/internal/task"
	"os/exec"
)

func Execute(task task.Task) error {
	cmd := exec.Command(task.Command, task.Args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil
	}

	if err = cmd.Start(); err != nil {
		return nil
	}

	scanner := bufio.NewScanner(stdOut)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return cmd.Wait()
}

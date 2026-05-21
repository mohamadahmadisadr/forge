package executor

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Execute(task string, args []string) {
	cmd := exec.Command(task, args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdOut)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

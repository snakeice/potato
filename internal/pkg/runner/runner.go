package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/snakeice/potato/internal/pkg/definitions"
)

func RunCommand(config *definitions.PotatoConfig, commandStr []string) {
	command := config.Shell

	if config.AlwaysSudo {
		command = append([]string{"sudo"}, command...)
	}

	if len(config.Shell) > 0 {
		command = append(command, strings.Join(commandStr, " "))
	} else {
		command = append(command, strings.Join(commandStr, " "))
	}

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Dir = "."

	fmt.Printf("Running %s\n", strings.Join(commandStr, " "))

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

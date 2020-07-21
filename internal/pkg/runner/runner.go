package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/atotto/clipboard"
	"github.com/snakeice/potato/internal/pkg/definitions"
)

func RunCommand(config *definitions.PotatoConfig, commandStr string) {
	err := clipboard.WriteAll(commandStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStarting %s\n", commandStr)
	splited := strings.Split((commandStr), " ")
	path, _ := exec.LookPath(splited[0])

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}
	proc, err := os.StartProcess(path, splited[1:], &pa)
	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Printf("<< Exited shell: %s\n", state.String())

}

func RunCommand_O(config *definitions.PotatoConfig, commandStr string) {
	err := clipboard.WriteAll(commandStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStarting %s\n", commandStr)
	splited := strings.Split((config.Shell), " ")
	splited = append(splited, commandStr)
	path, _ := exec.LookPath(splited[0])
	cmd := exec.Command(path, splited[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Dir = "."

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

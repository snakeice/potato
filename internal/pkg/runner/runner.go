package runner

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/snakeice/potato/internal/pkg/definitions"
)

func RunCommand(config *definitions.PotatoConfig, commandStr string) {
	splited := strings.Split((config.Shell + " " + commandStr), " ")
	cmd := exec.Command(splited[0], splited[0:]...)
	log.Printf("%v", cmd.Args)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

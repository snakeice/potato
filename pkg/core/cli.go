package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/snakeice/potato/internal/pkg/configuration"
	"github.com/snakeice/potato/internal/pkg/parser"
	"github.com/snakeice/potato/internal/pkg/runner"
)

func RunPotato() {
	config := configuration.LoadConfiguration()
	loadFunctions(config)
	if len(os.Args) <= 1 {
		fmt.Println("You need to pass a command")

		config.Commands["help"].Fn(config)
		os.Exit(1)
	}

	commandName := strings.Join(os.Args[1:], " ")

	if command, ok := config.Commands[commandName]; !ok {
		fmt.Printf("Command %s not found\n", commandName)
		os.Exit(1)
	} else {
		if command.Fn != nil {
			command.Fn(config)
			os.Exit(0)
		} else {
			parsedCommand := parser.MakeCommand(&command)
			cmd := strings.Split(strings.Join(parsedCommand, " "), " ")
			runner.RunCommand(config, cmd)
		}
	}

}

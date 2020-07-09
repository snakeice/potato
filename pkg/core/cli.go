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
	if len(os.Args) <= 1 {
		fmt.Printf("Use '%s <command>'.\n\n", os.Args[0])
		os.Exit(1)
	}

	commandName := strings.Join(os.Args[1:], " ")
	config := configuration.LoadConfiguration()
	loadFunctions()

	if call, ok := defaultFuncs[commandName]; ok {
		call(config)
		return
	}

	if command, ok := config.Commands[commandName]; !ok {
		fmt.Printf("Command %s not found\n", commandName)
		os.Exit(1)
	} else {
		parsedCommand := parser.MakeCommand(&command)
		runner.RunCommand(config, parsedCommand)
	}

}

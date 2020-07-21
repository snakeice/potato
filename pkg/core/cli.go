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
	loadFunctions()

	if len(os.Args) <= 1 {
		fmt.Printf("Use '%s <command>'.\n", os.Args[0])
		fmt.Println("Commands:")
		for name := range config.Commands {
			fmt.Printf("\t%s\n", name)
		}
		os.Exit(1)
	}

	commandName := strings.Join(os.Args[1:], " ")

	if fn, ok := defaultFuncs[commandName]; ok {
		fn(config)
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

package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/snakeice/potato/internal/pkg/definitions"
	"github.com/snakeice/potato/internal/pkg/runner"
)

func loadFunctions(config *definitions.PotatoConfig) {
	config.Commands["list"] = definitions.Command{
		Description: "List all commands",
		Fn: func(config *definitions.PotatoConfig) {
			for name, def := range config.Commands {
				fmt.Printf("\t%s - %s\n", name, def.Description)
			}
		},
	}

	config.Commands["help"] = definitions.Command{
		Description: "Show this help",
		Fn: func(config *definitions.PotatoConfig) {
			fmt.Printf("Use '%s <command>'\n", os.Args[0])

			println("Commands:")
			for name, def := range config.Commands {
				fmt.Printf("\t%s - %s\n", name, def.Description)
			}
		},
	}

	config.Commands["edit"] = definitions.Command{
		Description: "Edit config file",
		Fn: func(config *definitions.PotatoConfig) {
			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "vim"
			}

			cmd := strings.Split(editor, " ")
			cmd = append(cmd, config.Path)

			runner.RunCommand(config, cmd)
		},
	}
}

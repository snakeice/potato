package core

import (
	"fmt"

	"github.com/snakeice/potato/internal/pkg/definitions"
	"github.com/snakeice/potato/internal/pkg/runner"
)

type baseFunc = func(config *definitions.PotatoConfig)

var defaultFuncs = make(map[string]baseFunc)

func loadFunctions() {
	defaultFuncs["list"] = func(config *definitions.PotatoConfig) {
		for name, def := range config.Commands {
			println(name, def.Description)
		}
	}

	defaultFuncs["edit"] = func(config *definitions.PotatoConfig) {
		runner.RunCommand(config, fmt.Sprintf("%s %s", config.Editor, config.Path))

	}
}
package parser

import (
	"github.com/charmbracelet/huh"
	"github.com/snakeice/potato/internal/pkg/definitions"
)

func fillParams(command *definitions.Command) map[string]string {
	result := make(map[string]string)

	fields := []huh.Field{}

	for name, param := range command.Parameters {
		var description string
		if len(param.Description) == 0 {
			description = name
		} else {
			description = param.Description
		}

		if len(param.Values) == 0 {
			fields = append(fields, huh.NewInput().
				Key(name).
				Title(description).
				Suggestions([]string{param.Default}))

		} else {
			var options []huh.Option[string]
			for key, value := range param.Values {
				options = append(options, huh.NewOption(key, value))
			}

			fields = append(fields, huh.NewSelect[string]().
				Key(name).
				Title(description).
				Options(options...))
		}
	}

	form := huh.NewForm(huh.NewGroup(fields...))

	err := form.Run()
	if err != nil {
		return nil
	}

	for name := range command.Parameters {
		result[name] = form.GetString(name)
	}

	return result
}

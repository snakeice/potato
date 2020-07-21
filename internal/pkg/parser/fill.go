package parser

import (
	"log"
	"reflect"

	"github.com/manifoldco/promptui"
	"github.com/snakeice/potato/internal/pkg/definitions"
)

func fillParams(command *definitions.Command) map[string]string {
	result := make(map[string]string)
	for name, param := range command.Parameters {
		var description string
		if len(param.Description) == 0 {
			description = name
		} else {
			description = param.Description
		}

		if len(param.Values) == 0 {
			prompt := promptui.Prompt{
				Label:   description,
				Default: param.Default,
			}

			if response, err := prompt.Run(); err != nil {
				log.Printf("Err: %v", err)
			} else {
				result[name] = response
			}

		} else {
			prompt := promptui.Select{
				Label: description,
				Items: reflect.ValueOf(param.Values).MapKeys(),
			}
			if _, response, err := prompt.Run(); err != nil {
				log.Printf("Err: %v", err)
			} else {
				result[name] = param.Values[response]
			}
		}

	}

	return result
}

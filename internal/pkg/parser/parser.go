package parser

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/snakeice/potato/internal/pkg/definitions"
)

func parseCommand(templateStr []string, params map[string]string) []string {
	var result []string
	for _, part := range templateStr {

		tpl, err := template.New("proc").Parse(part)
		if err != nil {
			fmt.Println("Fail on load template")
			log.Panic(err)
		}
		var tplResult bytes.Buffer
		if err := tpl.Execute(&tplResult, params); err != nil {
			fmt.Println("Fail on parse template")
			log.Panic(err)
		}
		result = append(result, tplResult.String())
	}
	return result
}

func MakeCommand(command *definitions.Command) []string {
	params := fillParams(command)
	return parseCommand(command.Template, params)
}

package main

import (
	"log"
	"os"
)

const PROJECT_NAME string = "kutil-script"

type subCommand struct {
	Name        string
	Fun         func([]string)
	Description string
}

var SubCommands []subCommand = []subCommand{
	newCmd(
		"help",
		help,
		"display help",
	),
}

func newCmd(name string, fun func([]string), description string) subCommand {
	return subCommand{
		Name:        name,
		Fun:         fun,
		Description: description,
	}
}

func main() {
	log.Print("start " + PROJECT_NAME)

	args := os.Args

	if len(args) == 0 {
		help(args)
		return
	}

	for _, command := range SubCommands {
		if args[0] == command.Name {
			command.Fun(args)
			return
		}
	}

	help(args)
}

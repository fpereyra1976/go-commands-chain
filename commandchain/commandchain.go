package commandchain

import (
	"log"
	"reflect"
)

type CommandChain struct {
	commands []Command
}

func (commandChain *CommandChain) Execute(ctx *Context) bool {
	for _, command := range commandChain.commands {
		log.Println("Ejecutando: ",reflect.TypeOf(command))
		if !command.Execute(ctx) {
			return false
		}
	}
	return true
}

func (commandChain *CommandChain) AddCommand(c Command) {
	commandChain.commands = append(commandChain.commands, c)
}

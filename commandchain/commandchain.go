package commandchain

import (
	"log"
	"reflect"
)

type CommandChain struct {
	commands []Runnable
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

func (commandChain *CommandChain) AddCommand(r Runnable) {
	commandChain.commands = append(commandChain.commands, r)
}

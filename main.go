package main

import (
	"fmt"
	"github.com/fpereyra1976/go-commands-chain/commandchain"
)

type Command1 struct{}
type Command2 struct{}
type Command3 struct{}

func (c1 *Command1) Execute(ctx *commandchain.Context) bool {
	return true
}

func (c2 *Command2) Execute(ctx *commandchain.Context) bool {
	return true
}

func (c3 *Command3) Execute(ctx *commandchain.Context) bool {
	return true
}

func main() {
	chain := commandchain.CommandChain{}
	chain2 := commandchain.CommandChain{}

	ctx := commandchain.Context{Route: "Cadena: "}

	c1 := Command1{}
	c2 := Command2{}
	c3 := Command3{}

	c4_1 := Command1{}
	c4_2 := Command2{}
	c4_3 := Command3{}

	chain2.AddCommand(&c4_1)
	chain2.AddCommand(&c4_2)
	chain2.AddCommand(&c4_3)

	chain.AddCommand(&c1)
	chain.AddCommand(&c2)
	chain.AddCommand(&c3)
	chain.AddCommand(&chain2)

	if chain.Execute(&ctx) {
		fmt.Println("Ejecución Exitosa")
	} else {
		fmt.Println("Ejecución Fallida")
	}
}

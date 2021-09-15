package main

import (
	"fmt"

	"github.com/fpereyra1976/go-commands-chain/commandchain"
)

/* Middle Mile */
type CheckLHCapacitiesCommand struct{}
type WriteLHCapacitiesCommand struct{}

func (c3 *CheckLHCapacitiesCommand) Execute(ctx *commandchain.Context) bool {
	fmt.Println("CheckLHCapacitiesCommand.Execute: ", *ctx.GetParameter("Route"))
	return true
}

func (c3 *WriteLHCapacitiesCommand) Execute(ctx *commandchain.Context) bool {
	fmt.Println("WriteLHCapacitiesCommand.Execute: ", *ctx.GetParameter("Route"))
	return true
}

/* Middle Mile */

/* Last Mile */
type CheckCapacitiesCommand struct{}
type WriteCapacitiesCommand struct{}

func (c3 *CheckCapacitiesCommand) Execute(ctx *commandchain.Context) bool {
	fmt.Println("CheckCapacitiesCommand.Execute: ", *ctx.GetParameter("Route"))
	return true
}

func (c3 *WriteCapacitiesCommand) Execute(ctx *commandchain.Context) bool {
	fmt.Println("WriteCapacitiesCommand.Execute: ", *ctx.GetParameter("Route"))
	return true
}

/* Last Mile */

func main() {
	decorateRouteChain := commandchain.CommandChain{}
	decorateLastMileChain := commandchain.CommandChain{}
	decorateMiddleMileChain := commandchain.CommandChain{}

	/* Middle Mile */
	checkLHCapacitiesCommand := CheckLHCapacitiesCommand{}
	writeLHCapacitiesCommand := WriteLHCapacitiesCommand{}

	decorateMiddleMileChain.AddCommand(&checkLHCapacitiesCommand)
	decorateMiddleMileChain.AddCommand(&writeLHCapacitiesCommand)

	/* Last Mile */
	checkCapacitiesCommand := CheckCapacitiesCommand{}
	writeCapacitiesCommand := WriteCapacitiesCommand{}

	decorateLastMileChain.AddCommand(&checkCapacitiesCommand)
	decorateLastMileChain.AddCommand(&writeCapacitiesCommand)

	/* Route */
	ctx := commandchain.Context{}

	var route = new(commandchain.BusinessParameters)
	*route = "Hola"
	ctx.AddParameter("Route", route)

	decorateRouteChain.AddCommand(&decorateMiddleMileChain)
	decorateRouteChain.AddCommand(&decorateLastMileChain)

	if decorateRouteChain.Execute(&ctx) {
		fmt.Println("Ejecución Exitosa")
	} else {
		fmt.Println("Ejecución Fallida")
	}
}

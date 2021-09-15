package commandchain

type Command struct{}

func (cmd *Command) Execute(ctx *Context) bool {
	return false
}

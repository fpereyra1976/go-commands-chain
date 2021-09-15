package commandchain

type Command interface {
	Execute(ctx *Context) bool
}
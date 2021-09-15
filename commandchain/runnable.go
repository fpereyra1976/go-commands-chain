package commandchain

type CommandContext map[string]string

type Runnable interface {
	Execute(ctx *Context) bool
}

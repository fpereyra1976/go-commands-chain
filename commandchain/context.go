package commandchain

type BusinessParameters interface{}

type Context struct {
	parameters map[string]*BusinessParameters
}

func (ctx *Context) AddParameter(name string, param *BusinessParameters) {
	if ctx.parameters == nil {
		ctx.parameters = make(map[string]*BusinessParameters)
	}
	ctx.parameters[name] = param
}

func (ctx *Context) GetParameter(name string) *BusinessParameters {
	if ctx.parameters != nil {
		return ctx.parameters[name]
	}
	return nil;
}

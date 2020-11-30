package interpreter

type CallFunc func(i *Interpreter, args []interface{}) interface{}

func funcDefine(funcName string, call CallFunc) Symbol {
	return &FunctionSymbol{
		Name:       funcName,
		Type:       NewBuiltinFunctionSymbol(),
		ScopeLevel: 1,
		Call:       call,
	}
}

var LenFunc = funcDefine("len", func(i *Interpreter, args []interface{}) interface{} {
	if len(args) != 1 {
		return 0
	}
	if v, ok := args[0].(string); ok {
		return len(v)
	}
	return 0
})

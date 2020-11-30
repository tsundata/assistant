package interpreter

func packageDefine(packageName string, calls map[string]CallFunc) []Symbol {
	var packageCalls []Symbol
	for funcName, call := range calls {
		packageCalls = append(packageCalls, &FunctionSymbol{
			Package:    packageName,
			Name:       funcName,
			ScopeLevel: 1,
			Call:       call,
		})
	}
	return packageCalls
}

var packages = map[string][]Symbol{
	"net": packageDefine("net", map[string]CallFunc{
		"get": func(i *Interpreter, args []interface{}) interface{} {
			return "net.get"
		},
	}),
	"message": packageDefine("message", map[string]CallFunc{
		"send": func(i *Interpreter, args []interface{}) interface{} {
			return "message.send"
		},
	}),
}

package interpreter

func packageDefine(packageName string, calls map[string]CallFunc) map[string]Symbol {
	packageCalls := make(map[string]Symbol)
	for funcName, call := range calls {
		packageCalls[funcName] = &FunctionSymbol{
			Package:    packageName,
			Name:       funcName,
			ScopeLevel: 1,
			Call:       call,
		}
	}
	return packageCalls
}

var packages = map[string]map[string]Symbol{
	// net
	"net": packageDefine("net", map[string]CallFunc{
		"get": func(i *Interpreter, args []interface{}) interface{} {
			return "net.get"
		},
	}),

	// message
	"msg": packageDefine("msg", map[string]CallFunc{
		"send": func(i *Interpreter, args []interface{}) interface{} {
			return "message.send" + args[0].(string)
		},
		"filter": func(i *Interpreter, args []interface{}) interface{} {
			return len(args[0].(string)) > 5
		},
		"reduce": func(i *Interpreter, args []interface{}) interface{} {
			return args[1].(string) + args[0].(string)
		},
	}),
}

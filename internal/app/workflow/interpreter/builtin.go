package interpreter

const BuiltinPackage = "builtin"

type CallFunc func(i *Interpreter, args []interface{}) interface{}

func funcDefine(funcName string, call CallFunc) Symbol {
	return &FunctionSymbol{
		Package:    BuiltinPackage,
		Name:       funcName,
		ScopeLevel: 1,
		Call:       call,
	}
}

var functions = map[string]Symbol{
	// len
	"len": funcDefine("len", func(i *Interpreter, args []interface{}) interface{} {
		if len(args) != 1 {
			return 0
		}
		if v, ok := args[0].(string); ok {
			return len(v)
		}
		return 0
	}),
}

var iteration = map[string]Symbol{
	// map
	"map": funcDefine("map", func(i *Interpreter, args []interface{}) interface{} {
		if len(args) != 2 {
			return nil
		}

		call := tfCall(args[1])
		var result []interface{}
		for _, item := range args[0].([]interface{}) {
			r := call(i, []interface{}{item})
			result = append(result, r)
		}

		return result
	}),

	// filter
	"filter": funcDefine("filter", func(i *Interpreter, args []interface{}) interface{} {
		if len(args) != 2 {
			return nil
		}

		call := tfCall(args[1])
		var result []interface{}
		for _, item := range args[0].([]interface{}) {
			r := call(i, []interface{}{item})
			if v, ok := r.(bool); ok && v {
				result = append(result, item)
			}
		}

		return result
	}),

	// reduce
	"reduce": funcDefine("reduce", func(i *Interpreter, args []interface{}) interface{} {
		if len(args) != 2 {
			return nil
		}

		call := tfCall(args[1])

		items := args[0].([]interface{})
		if len(items) == 0 {
			return nil
		}
		if len(items) == 1 {
			return items[0]
		}
		var ins [2]interface{}
		ins[0] = items[0]
		ins[1] = items[1]
		result := call(i, []interface{}{ins[0], ins[1]})
		for index := 2; index < len(items); index++ {
			ins[0] = result
			ins[1] = items[index]
			result = call(i, []interface{}{ins[0], ins[1]})
		}

		return result
	}),
}

func tfCall(arg interface{}) CallFunc {
	ref := arg.(*FunctionRef)

	var call CallFunc
	if ref.PackageName == "" || ref.PackageName == BuiltinPackage {
		call = functions[ref.FuncName].(*FunctionSymbol).Call
	} else {
		call = packages[ref.PackageName][ref.FuncName].(*FunctionSymbol).Call
	}

	return call
}

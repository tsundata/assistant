package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"strings"
)

type Opcoder interface {
	Run(ctx *inside.Context, params []interface{}) (interface{}, error)
}

func RunOpcode(ctx *inside.Context, name string, params []interface{}) (interface{}, error) {
	var o Opcoder
	switch strings.ToLower(name) {
	case "get":
		o = NewGet()
	case "count":
		o = NewCount()
	case "send":
		o = NewSend()
	case "task":
		o = NewTask()
	case "debug":
		o = NewDebug()
	case "json":
		o = NewJson()
	case "if":
		o = NewIf()
	case "else":
		o = NewElse()
	case "set":
		o = NewSet()
	default:
		return nil, errors.New("not opcode")
	}
	return o.Run(ctx, params)
}

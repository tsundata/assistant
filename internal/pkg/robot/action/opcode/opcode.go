package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"sort"
	"strings"
)

const (
	TypeOp = iota
	TypeCond
	TypeAsync
)

type Opcoder interface {
	Type() int
	Doc() string
	Run(ctx context.Context, comp *inside.Component, params []interface{}) (interface{}, error)
}

var opcodes = map[string]Opcoder{
	"cron":    NewCron(),
	"webhook": NewWebhook(),
	"if":      NewIf(),
	"else":    NewElse(),
	"get":     NewGet(),
	"count":   NewCount(),
	"echo":    NewEcho(),
	"task":    NewTask(),
	"debug":   NewDebug(),
	"json":    NewJson(),
	"set":     NewSet(),
	"status":  NewStatus(),
	"message": NewMessage(),
	"query":   NewQuery(),
	"dedupe":  NewDedupe(),
	"secret":  NewSecret(),
	"env":     NewEnv(),
}

func NewOpcode(name string) Opcoder {
	if op, ok := opcodes[strings.ToLower(name)]; ok {
		return op
	}
	return nil
}

func Docs() []string {
	keys := make([]string, 0, len(opcodes))
	for k := range opcodes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var res []string
	for _, k := range keys {
		res = append(res, opcodes[k].Doc())
	}

	return res
}

func Doc(op string) string {
	if o, ok := opcodes[strings.ToLower(op)]; ok {
		return o.Doc()
	}
	return ""
}

package expr

import (
	"context"
	"github.com/antonmedv/expr"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strings"
)

// Expr https://github.com/antonmedv/expr/blob/master/docs/Language-Definition.md
type Expr struct{}

func (a Expr) Run(_ context.Context, _ *bot.Controller, param []util.Value, input bot.PluginValue) (bot.PluginValue, error) {
	var in interface{}
	if len(param) > 0 {
		if code, ok := param[0].String(); ok {
			program, err := expr.Compile(code, expr.Env(Env{}))
			if err != nil {
				return bot.PluginValue{}, err
			}
			output, err := expr.Run(program, Env{Value: input.Value})
			if err != nil {
				return bot.PluginValue{}, err
			}
			in = output
			if s, ok := output.(string); ok {
				input.Value = s
			}
		}
	}
	input.Stack = append(input.Stack, in)
	return input, nil
}

func (a Expr) Name() string {
	return "expr"
}

type Env struct {
	Value string
}

func (Env) Trim(str string) string {
	return strings.TrimSpace(str)
}

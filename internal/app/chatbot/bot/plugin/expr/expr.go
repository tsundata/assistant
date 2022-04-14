package expr

import (
	"context"
	"github.com/antonmedv/expr"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"strings"
)

// Expr https://github.com/antonmedv/expr/blob/master/docs/Language-Definition.md
type Expr struct {
	Next bot.PluginHandler
}

func (a Expr) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	var in interface{}
	params := bot.Param(ctrl, a)
	if len(params) > 0 {
		if code, ok := params[0].(string); ok {
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
	input.Stack[a.Name()] = in
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
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

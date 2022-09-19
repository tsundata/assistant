package expr

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"testing"
)

func TestExpr(t *testing.T) {
	p := Expr{}

	type args struct {
		input  string
		code   string
		output string
		stack  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{"test", "len(Value)", "test", 4}, false},
		{"case2", args{"test  ", "Trim(Value)", "test", "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := bot.PluginValue{Value: tt.args.input, Stack: []interface{}{}}
			ctrl := &bot.Controller{}
			params := []util.Value{
				util.Variable(tt.args.code),
			}
			output, err := p.Run(context.Background(), ctrl, params, input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestExpr error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			input.Value = tt.args.output
			input.Stack = []interface{}{tt.args.stack}
			assert.Equal(t, output.Stack[0], tt.args.stack)
			assert.Equal(t, input, output)
		})
	}
}

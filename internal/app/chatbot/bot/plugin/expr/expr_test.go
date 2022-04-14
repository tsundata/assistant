package expr

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestExpr(t *testing.T) {
	p := Expr{
		Next: end.End{},
	}

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
			input := bot.PluginValue{Value: tt.args.input, Stack: make(map[string]interface{})}
			ctrl := bot.MockController(map[string][]interface{}{
				"expr": {tt.args.code},
			})
			output, err := p.Run(context.Background(), ctrl, input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestExpr error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			input.Value = tt.args.output
			assert.Equal(t, output.Stack[p.Name()], tt.args.stack)
			assert.Equal(t, input, output)
		})
	}
}

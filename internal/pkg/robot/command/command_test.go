package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strconv"
	"testing"
)

func TestRegexRule(t *testing.T) {
	var testRules = []Rule{
		{
			Define: `test`,
			Help:   `Test info`,
			Parse: func(ctx context.Context, comp component.Component, tokens []*Token) []pb.MsgPayload {
				return []pb.MsgPayload{pb.TextMsg{Text: "test"}}
			},
		},
		{
			Define: `todo [string]`,
			Help:   `todo something`,
			Parse: func(ctx context.Context, comp component.Component, tokens []*Token) []pb.MsgPayload {
				text, _ := tokens[1].Value.String()
				return []pb.MsgPayload{pb.TextMsg{Text: text}}
			},
		},
		{
			Define: `add [number] [number]`,
			Help:   `Addition`,
			Parse: func(ctx context.Context, comp component.Component, tokens []*Token) []pb.MsgPayload {
				tt1, _ := tokens[1].Value.Int64()
				tt2, _ := tokens[2].Value.Int64()
				return []pb.MsgPayload{pb.TextMsg{Text: strconv.Itoa(int(tt1 + tt2))}}
			},
		},
	}

	b := New(testRules)

	out, err := b.ProcessCommand(context.Background(), nil, "test")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out, pb.TextMsg{Text: "test"})

	out2, err := b.ProcessCommand(context.Background(), nil, "add 1 2")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out2, pb.TextMsg{Text: "3"})

	out3, err := b.ProcessCommand(context.Background(), nil, "help")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, out3, 0)

	help := b.Help("help")
	assert.True(t, len(help) > 0)

	out4, err := b.ProcessCommand(context.Background(), nil, `todo "a b c"`)
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out4, pb.TextMsg{Text: "a b c"})
}

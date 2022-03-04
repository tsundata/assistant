package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"strconv"
	"testing"
)

func TestRegexRule(t *testing.T) {
	var testRules = []Rule{
		{
			Define: `test`,
			Help:   `Test info`,
			Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
				return []string{"test"}
			},
		},
		{
			Define: `todo [string]`,
			Help:   `todo something`,
			Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
				if len(tokens) != 2 {
					return []string{"error args"}
				}

				return []string{
					tokens[1].Value,
				}
			},
		},
		{
			Define: `add [number] [number]`,
			Help:   `Addition`,
			Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
				if len(tokens) != 3 {
					return []string{"error args"}
				}

				tt1, err := strconv.ParseInt(tokens[1].Value, 10, 64)
				if err != nil {
					return []string{"error call: " + err.Error()}
				}

				tt2, err := strconv.ParseInt(tokens[2].Value, 10, 64)
				if err != nil {
					return []string{"error call: " + err.Error()}
				}

				return []string{
					strconv.Itoa(int(tt1 + tt2)),
				}
			},
		},
	}

	b := New(testRules)

	out, err := b.ParseCommand(context.Background(), nil, "test")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out, "test")

	out2, err := b.ParseCommand(context.Background(), nil, "add 1 2")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out2, "3")

	out3, err := b.ParseCommand(context.Background(), nil, "help")
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, out3, 0)

	help := b.Help("help")
	assert.True(t, len(help) > 0)

	out4, err := b.ParseCommand(context.Background(), nil, `todo "a b c"`)
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out4, "a b c")
}

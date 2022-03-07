package command

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestRegexRule(t *testing.T) {
	var testRules = []Rule{
		{
			Define: `test`,
			Help:   `Test info`,
			Parse: func(ctx context.Context, comp Component, tokens []*Token) []string {
				return []string{"test"}
			},
		},
		{
			Define: `todo [string]`,
			Help:   `todo something`,
			Parse: func(ctx context.Context, comp Component, tokens []*Token) []string {
				return []string{
					tokens[1].Value.(string),
				}
			},
		},
		{
			Define: `add [number] [number]`,
			Help:   `Addition`,
			Parse: func(ctx context.Context, comp Component, tokens []*Token) []string {
				tt1 := tokens[1].Value.(int64)
				tt2 := tokens[2].Value.(int64)

				return []string{
					strconv.Itoa(int(tt1 + tt2)),
				}
			},
		},
	}

	b := New(testRules)

	out, err := b.ProcessCommand(context.Background(), nil, "test")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out, "test")

	out2, err := b.ProcessCommand(context.Background(), nil, "add 1 2")
	if err != nil {
		t.Fatal(err)
	}
	require.Contains(t, out2, "3")

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
	require.Contains(t, out4, "a b c")
}

package rule

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
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

	var options = []rulebot.Option{
		rulebot.RegisterRuleset(New(testRules)),
	}

	b := rulebot.New(nil)
	b.SetOptions(options...)

	res := b.Process(context.Background(), "test")
	require.Contains(t, res.MessageProviderOut(), "test")

	res2 := b.Process(context.Background(), "add 1 2")
	require.Contains(t, res2.MessageProviderOut(), "3")

	res3 := b.Process(context.Background(), "help")
	require.Len(t, res3.MessageProviderOut(), 1)

	res4 := b.Process(context.Background(), `todo "a b c"`)
	require.Contains(t, res4.MessageProviderOut(), "a b c")
}

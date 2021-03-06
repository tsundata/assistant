package rule

import (
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"strconv"
	"testing"
)

func TestRegexRule(t *testing.T) {
	var testRules = []Rule{
		{
			Regex:       `test`,
			HelpMessage: `Test info`,
			ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
				return []string{"test"}
			},
		},
		{
			Regex:       `add\s+(\d+)\s+(\d+)`,
			HelpMessage: `Addition`,
			ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
				if len(args) != 3 {
					return []string{"error args"}
				}

				tt1, err := strconv.ParseInt(args[1], 10, 64)
				if err != nil {
					return []string{"error call: " + err.Error()}
				}

				tt2, err := strconv.ParseInt(args[2], 10, 64)
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

	res := b.Process("test")
	require.Contains(t, res.MessageProviderOut(), "test")

	res2 := b.Process("add 1 2")
	require.Contains(t, res2.MessageProviderOut(), "3")

	res3 := b.Process("help")
	require.Len(t, res3.MessageProviderOut(), 1)
}

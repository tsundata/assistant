package plugins

import (
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/cron"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/regex"
	"github.com/tsundata/assistant/internal/pkg/model"
	"log"
	"testing"
)

func TestRunPlugin(t *testing.T) {
	regexRules = []regex.Rule{
		{
			Regex:       `godoc (.*)`,
			HelpMessage: `search godoc.org and return the first result`,
			ParseMessage: func(bot *bot.Bot, s string, args []string) []string {
				return []string{
					args[1] + " ..... doc .....",
					args[1] + " ..... doc .....",
					args[1] + " ..... doc .....",
					args[1] + " ..... doc .....",
					args[1] + " ..... doc .....",
					args[1] + " ..... doc .....",
				}
			},
		},
		{
			Regex:       `{{ .RobotName }} hi (.*)`,
			HelpMessage: `Demo plugin`,
			ParseMessage: func(bot *bot.Bot, s string, args []string) []string {
				return []string{
					args[1] + " ..... hi .....",
				}
			},
		},
	}
	Options = []bot.Option{
		bot.RegisterRuleset(regex.New(regexRules)),
		bot.RegisterRuleset(cron.New(cronRules)),
	}

	b := bot.New("test", nil, nil, Options...)

	out := b.Process(model.Event{
		Data: model.EventData{Message: model.Message{
			Text: "test hi abc",
		}},
	}).MessageProviderOut()
	log.Println(out)

	out2 := b.Process(model.Event{
		Data: model.EventData{Message: model.Message{
			Text: "test help",
		}},
	}).MessageProviderOut()
	log.Println(out2)
}

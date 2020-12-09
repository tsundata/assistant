package plugins

import (
	"context"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/cron"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/regex"
	"github.com/tsundata/assistant/internal/pkg/model"
	"log"
	"time"
)

var regexRules = []regex.Rule{
	{
		Regex:       `demo (.*)`,
		HelpMessage: `demo plugin`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			return []string{
				args[1] + "Hello world " + time.Now().String(),
			}
		},
	},
	{
		Regex:       `subs list`,
		HelpMessage: `List subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			var reply []string
			err := b.SubClient.Call(context.Background(), "List", nil, &reply)
			if err != nil {
				return []string{"error call"}
			}

			return reply
		},
	},
	{
		Regex:       `subs open (.*)`,
		HelpMessage: `Open subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}
			var reply bool
			err := b.SubClient.Call(context.Background(), "Open", &args[1], &reply)
			if err != nil {
				return []string{"error call"}
			}
			if reply {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex:       `subs close (.*)`,
		HelpMessage: `Close subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}
			var reply bool
			err := b.SubClient.Call(context.Background(), "Close", &args[1], &reply)
			if err != nil {
				return []string{"error call"}
			}
			if reply {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
}

var cronRules = map[string]cron.Rule{
	"heartbeat": {
		"0 0 * * *",
		func() []model.Event {
			log.Println("cron " + time.Now().String())
			return []model.Event{
				{
					Data: model.EventData{Message: model.Message{
						Text: "Plugin Cron Heartbeat: " + time.Now().String(),
					}},
				},
			}
		},
	},
}

var Options = []bot.Option{
	bot.RegisterRuleset(regex.New(regexRules)),
	bot.RegisterRuleset(cron.New(cronRules)),
}

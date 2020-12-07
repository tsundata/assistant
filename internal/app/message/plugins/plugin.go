package plugins

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/cron"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/regex"
	"github.com/tsundata/assistant/internal/pkg/model"
	"time"
)

var regexRules = []regex.Rule{
	{
		Regex:       `demo (.*)`,
		HelpMessage: `demo plugin`,
		ParseMessage: func(s string, args []string) []string {
			return []string{
				args[1] + "Hello world " + time.Now().String(),
			}
		},
	},
}
var cronRules = map[string]cron.Rule{
	"heartbeat": {
		"0 0 * * *",
		func() []model.Event {
			fmt.Println("cron " + time.Now().String())
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

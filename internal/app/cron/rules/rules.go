package rules

import (
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"time"
)

var rules = map[string]Rule{
	"heartbeat": {
		When: "0 0 * * *",
		Action: func() []model.Event {
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

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}

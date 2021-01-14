package rules

import (
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"time"
)

var rules = map[string]Rule{
	"heartbeat": {
		When: "0 0 * * *",
		Action: func() []model.Message {
			return []model.Message{
				{
					Text: "Plugin Cron Heartbeat: " + time.Now().String(),
				},
			}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}

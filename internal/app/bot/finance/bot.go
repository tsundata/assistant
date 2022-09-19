package finance

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var metadata = bot.Metadata{
	Name:       "Finance",
	Identifier: enum.FinanceBot,
	Detail:     "",
	Avatar:     "",
}

var workflowRules bot.WorkflowRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, nil, nil, workflowRules, commandRules, nil, nil, nil)
	if err != nil {
		panic(err)
	}
}

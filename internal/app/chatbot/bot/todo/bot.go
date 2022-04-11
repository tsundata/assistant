package todo

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var metadata = bot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
}

var setting = []bot.FieldItem{
	{
		Key:      "report",
		Type:     bot.FieldItemTypeBool,
		Required: false,
		Value:    false,
	},
	{
		Key:      "time",
		Type:     bot.FieldItemTypeString,
		Required: true,
		Value:    "",
	},
}

var workflowRules bot.WorkflowRule

var actionRules []bot.ActionRule

var formRules []bot.FormRule

var tagRules []bot.TagRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules, actionRules, formRules, tagRules)
	if err != nil {
		panic(err)
	}
}

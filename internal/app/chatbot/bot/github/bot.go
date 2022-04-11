package github

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
)

var metadata = bot.Metadata{
	Name:       "Github",
	Identifier: enum.GithubBot,
	Detail:     "",
	Avatar:     "",
}

var setting []bot.FieldItem

var commandRules []command.Rule

var workflowRules bot.WorkflowRule

var actionRules []bot.ActionRule

var formRules []bot.FormRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules, actionRules, formRules, tagRules)
	if err != nil {
		panic(err)
	}
}

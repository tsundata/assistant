package org

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var metadata = bot.Metadata{
	Name:       "Org",
	Identifier: enum.OrgBot,
	Detail:     "",
	Avatar:     "",
}

var workflowRules []bot.PluginRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, nil, workflowRules, commandRules, nil, nil, nil)
	if err != nil {
		panic(err)
	}
}

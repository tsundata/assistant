package org

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var Metadata = bot.Metadata{
	Name:       "Org",
	Identifier: enum.OrgBot,
	Detail:     "",
	Avatar:     "",
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(Metadata, nil, nil)
	if err != nil {
		panic(err)
	}
}
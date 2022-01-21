package finance

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var Metadata = bot.Metadata{
	Name:       "Finance",
	Identifier: enum.FinanceBot,
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
package robot

import (
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/org"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var BotMap = map[string]*bot.Bot{
	system.Bot.Metadata.Identifier:  system.Bot,
	todo.Bot.Metadata.Identifier:    todo.Bot,
	org.Bot.Metadata.Identifier:     org.Bot,
	finance.Bot.Metadata.Identifier: finance.Bot,
}

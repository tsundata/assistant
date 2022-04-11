package robot

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/github"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/org"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var BotMap = map[string]*bot.Bot{
	enum.SystemBot:  system.Bot,
	enum.TodoBot:    todo.Bot,
	enum.OrgBot:     org.Bot,
	enum.FinanceBot: finance.Bot,
	enum.GithubBot:  github.Bot,
}

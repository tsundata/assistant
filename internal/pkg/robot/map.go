package robot

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/github"
	"github.com/tsundata/assistant/internal/app/bot/okr"
	"github.com/tsundata/assistant/internal/app/bot/system"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var BotMap = map[string]*bot.Bot{
	enum.SystemBot:  system.Bot,
	enum.TodoBot:    todo.Bot,
	enum.OkrBot:     okr.Bot,
	enum.FinanceBot: finance.Bot,
	enum.GithubBot:  github.Bot,
}

package app

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/github"
	"github.com/tsundata/assistant/internal/app/bot/okr"
	"github.com/tsundata/assistant/internal/app/bot/system"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/app/cron/rule"
	"github.com/tsundata/assistant/internal/app/listener"
	"github.com/tsundata/assistant/internal/app/repository"
	"github.com/tsundata/assistant/internal/app/service"
	"github.com/tsundata/assistant/internal/app/spider/crawler"
	"github.com/tsundata/assistant/internal/app/work"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
)

func NewApp(c *config.AppConfig, logger log.Logger, hs *http.Server, bus event.Bus, rdb *redis.Client, bot *rulebot.RuleBot, q *machinery.Server,
	comp component.Component,
	message service.MessageSvcClient,
	middle service.MiddleSvcClient,
	user service.UserSvcClient,
	chatbot service.ChatbotSvcClient,
	storage pb.StorageSvcClient,
	messageRepo repository.MessageRepository,
	userRepo repository.UserRepository,
	chatbotRepo repository.ChatbotRepository) (*app.Application, error) {

	// event bus register
	err := listener.RegisterEventHandler(bus, logger, c, rdb, bot, comp, message, middle, chatbot, storage, messageRepo, userRepo, chatbotRepo)
	if err != nil {
		return nil, err
	}

	// bots register
	err = robot.RegisterBot(context.Background(), bus, comp, system.Bot, todo.Bot, okr.Bot, finance.Bot, github.Bot)
	if err != nil {
		return nil, err
	}

	// cron
	go func() {
		// Delayed loading
		//time.Sleep(1 * time.Minute)
		// load rule
		bot.SetOptions(rule.Options...)
		logger.Info("start cron rule bot")
	}()

	// spider
	go func() {
		s := crawler.New()
		s.SetService(c, rdb, bus, logger, middle, message, user)
		err := s.LoadRule()
		if err != nil {
			logger.Error(err)
			return
		}
		s.Daemon()
	}()

	// worker
	go func() {
		workflowTask := work.NewWorkflowTask(bus, message, chatbot)
		err = q.RegisterTasks(map[string]interface{}{
			enum.WorkflowRunTask: workflowTask.Run,
		})
		if err != nil {
			logger.Error(err)
			return
		}

		worker := q.NewWorker(c.Name, 0)
		worker.SetErrorHandler(func(err error) {
			logger.Error(err)
		})
		err = worker.Launch()
		if err != nil {
			logger.Error(err)
			return
		}
	}()

	// http server
	a, err := app.New(c, logger, app.HTTPServerOption(hs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)

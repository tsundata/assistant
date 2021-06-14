package worker

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/worker/task"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, server *machinery.Server, client *rpc.Client) (*app.Application, error) {
	a, err := app.New(c, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// worker
	go func() {
		workflowTask := task.NewWorkflowTask(bus, client)
		echoTask := task.NewEchoTask(bus, client)
		err = server.RegisterTasks(map[string]interface{}{
			"run":  workflowTask.Run,
			"echo": echoTask.Echo,
		})
		if err != nil {
			logger.Error(err)
			return
		}

		worker := server.NewWorker(c.Name, 0)
		worker.SetErrorHandler(func(err error) {
			logger.Error(err)
		})
		err = worker.Launch()
		if err != nil {
			logger.Error(err)
			return
		}
	}()

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)

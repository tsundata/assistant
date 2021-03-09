package worker

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/worker/tasks"
	"github.com/tsundata/assistant/internal/pkg/app"
	"go.uber.org/zap"
)

func NewApp(name string, logger *zap.Logger, server *machinery.Server, msgClient pb.MessageClient) (*app.Application, error) {
	a, err := app.New(name, logger)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// worker
	go func() {
		workflowTask := tasks.NewWorkflowTask(msgClient)
		err = server.RegisterTasks(map[string]interface{}{
			"run": workflowTask.Run,
		})
		if err != nil {
			logger.Error(err.Error())
			return
		}

		worker := server.NewWorker(name, 0)
		worker.SetErrorHandler(func(err error) {
			logger.Error(err.Error())
		})
		err = worker.Launch()
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}()

	return a, nil
}

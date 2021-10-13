package task

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/task/work"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(
	c *config.AppConfig,
	bus event.Bus,
	logger log.Logger,
	rs *rpc.Server,
	q *machinery.Server,
	message pb.MessageSvcClient,
	workflow pb.WorkflowSvcClient) (*app.Application, error) {

	a, err := app.New(c, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	// worker
	go func() {
		workflowTask := work.NewWorkflowTask(bus, message, workflow)
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

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)

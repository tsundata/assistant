package worker

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/worker/tasks"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"log"
	"os"
)

func NewApp(logger *logger.Logger, server *machinery.Server, mq *amqp.Connection, msgClient pb.MessageClient, wfClient pb.WorkflowClient) (*app.Application, error) {
	name := os.Getenv("APP_NAME")

	a, err := app.New(name, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// worker
	go func() {
		workflowTask := tasks.NewWorkflowTask(msgClient, wfClient)
		echoTask := tasks.NewEchoTask(msgClient, wfClient)
		err = server.RegisterTasks(map[string]interface{}{
			"run":  workflowTask.Run,
			"echo": echoTask.Echo,
		})
		if err != nil {
			logger.Error(err)
			return
		}

		worker := server.NewWorker(name, 0)
		worker.SetErrorHandler(func(err error) {
			logger.Error(err)
		})
		err = worker.Launch()
		if err != nil {
			logger.Error(err)
			return
		}
	}()

	// mq
	go func() {
		ch, err := mq.Channel()
		if err != nil {
			logger.Error(err)
			return
		}
		q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
		if err != nil {
			logger.Error(err)
			return
		}

		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			logger.Error(err)
			return
		}

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()
	}()

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)

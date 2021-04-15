package message

import (
	"errors"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/rules"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"google.golang.org/grpc"
	"os"
)

type Options struct {
	Name    string
	Webhook string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	o.Name = os.Getenv("APP_NAME")

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *logger.Logger, rs *rpc.Server, db *sqlx.DB, mq *amqp.Connection,
	subClient pb.SubscribeClient, midClient pb.MiddleClient, msgClient pb.MessageClient,
	taskClient pb.TaskClient, wfClient pb.WorkflowClient) (*app.Application, error) {

	b := rulebot.New(nil, mq, subClient, midClient, msgClient, wfClient, taskClient, rules.Options...)

	message := service.NewManage(db, logger, b, o.Webhook, wfClient, msgClient, midClient)
	err := rs.Register(func(s *grpc.Server) error {
		pb.RegisterMessageServer(s, message)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)

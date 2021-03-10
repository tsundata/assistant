package message

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"google.golang.org/grpc"
)

type Options struct {
	Webhook string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(name string, o *Options, logger *logger.Logger, rs *rpc.Server, db *sqlx.DB, b *rulebot.RuleBot, wfClient pb.WorkflowClient) (*app.Application, error) {
	message := service.NewManage(db, logger, b, o.Webhook, wfClient)
	err := rs.Register(func(s *grpc.Server) error {
		pb.RegisterMessageServer(s, message)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(name, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

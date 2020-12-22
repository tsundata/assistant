package message

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/bot"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Options struct {
	Name    string
	webhook string
	db      *gorm.DB
	logger  *zap.Logger
}

func NewOptions(v *viper.Viper, db *gorm.DB, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	o.db = db
	o.logger = logger

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	slack := v.GetStringMapString("slack")
	o.webhook = slack["webhook"]

	return o, err
}

func NewApp(o *Options, rs *rpc.Server, b *bot.Bot) (*app.Application, error) {
	message := service.NewManage(o.db, o.logger, b, o.webhook)
	err := rs.Register(func(s *grpc.Server) error {
		pb.RegisterMessageServer(s, message)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, o.logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

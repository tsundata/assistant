package workflow

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, rs *rpc.Server, etcd *clientv3.Client, midClient pb.MiddleClient, msgClient pb.MessageClient) (*app.Application, error) {
	// service
	subscribe := service.NewWorkflow(etcd, midClient, msgClient)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterWorkflowServer(gs, subscribe)
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

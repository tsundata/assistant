package service

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(ps *Task) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterTaskServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewTask, CreateInitServerFn)

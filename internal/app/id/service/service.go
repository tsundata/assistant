package service

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(ps *Id) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterIdSvcServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewId, CreateInitServerFn)

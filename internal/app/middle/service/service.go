package service

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(ps *Middle) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterMiddleSvcServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewMiddle, CreateInitServerFn)

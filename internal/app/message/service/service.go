package service

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(ps *Message) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterMessageServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewMessage, CreateInitServerFn)

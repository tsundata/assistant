package service

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(ps *Workflow) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterWorkflowServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewWorkflow, CreateInitServerFn)

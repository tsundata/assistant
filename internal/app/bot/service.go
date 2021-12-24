package bot

import (
	"github.com/tsundata/assistant/api/pb"
	service2 "github.com/tsundata/assistant/internal/app/bot/finance/service"
	service3 "github.com/tsundata/assistant/internal/app/bot/org/service"
	"github.com/tsundata/assistant/internal/app/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func CreateInitServerFn(todo *service.Todo, finance *service2.Finance, org *service3.Org) rpc.InitServer {
	return func(s *grpc.Server) {
		pb.RegisterTodoSvcServer(s, todo)
		pb.RegisterFinanceSvcServer(s, finance)
		pb.RegisterOrgSvcServer(s, org)
	}
}

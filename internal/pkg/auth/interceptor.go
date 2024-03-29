package auth

import (
	"context"
	"log"

	"github.com/tsundata/assistant/internal/pkg/transport/rpc/exception"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"google.golang.org/grpc"
)

var ignoreFullMethod = []string{
	// basic
	"/grpc.health.v1.Health/Check",
	"/grpc.health.v1.Health/Watch",
	"/pb.IdSvc/GetGlobalId",
	"/pb.UserSvc/Login",
	"/pb.UserSvc/Authorization",
	// business
	"/pb.ChatbotSvc/Register",
	"/pb.MiddleSvc/GetCronStatus",
	"/pb.MiddleSvc/RegisterSubscribe",
	"/pb.MiddleSvc/GetSubscribeStatus",
	"/pb.MiddleSvc/GetPage",
	"/pb.MiddleSvc/RegisterCron",
	"/pb.ChatbotSvc/CronTrigger",
	"/pb.ChatbotSvc/WebhookTrigger",
	"/pb.StorageSvc/AbsolutePath",
	"/pb.MessageSvc/GetById",
	"/pb.StorageSvc/AbsolutePath",
	"/pb.UserSvc/GetUsers",
	"/pb.MiddleSvc/CreatePage",
	"/pb.ChatbotSvc/WatchTrigger",
	"/pb.MiddleSvc/CollectMetadata",
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !util.In(ignoreFullMethod, info.FullMethod) {
			if _, ok := md.FromIncoming(ctx); !ok {
				log.Println("Unauthenticated", info.FullMethod)
				return nil, exception.ErrGrpcUnauthenticated
			}
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !util.In(ignoreFullMethod, info.FullMethod) {
			if _, ok := md.FromIncoming(ss.Context()); !ok {
				log.Println("Unauthenticated", info.FullMethod)
				return exception.ErrGrpcUnauthenticated
			}
		}
		return handler(srv, ss)
	}
}

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(md.TraceContext(ctx), method, req, reply, cc)
	}
}

func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(md.TraceContext(ctx), desc, cc, method, opts...)
	}
}

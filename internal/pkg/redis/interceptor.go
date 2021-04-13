package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"time"
)

func StatsUnaryServerInterceptor(rdb *redis.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)

		// stats count
		rdb.Incr(ctx, fmt.Sprintf("stats:count:%s", info.FullMethod))

		// stats month
		now := time.Now()
		rdb.SetBit(ctx, fmt.Sprintf("stats:month:%s", now.Format("2006:01")), int64(now.Day()), 1)

		return resp, err
	}
}

func StatsStreamServerInterceptor(rdb *redis.Client) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		ctx := context.Background()

		// stats count
		rdb.Incr(ctx, fmt.Sprintf("stats:count:%s", info.FullMethod))

		// stats month
		now := time.Now()
		rdb.SetBit(ctx, fmt.Sprintf("stats:month:%s", now.Format("2006:01")), int64(now.Day()), 1)

		return err
	}
}

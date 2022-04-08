package md

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/enum"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func Outgoing(c *fiber.Ctx) context.Context {
	md := metadata.Pairs()
	idVal := c.Locals(enum.AuthKey)
	if id, ok := idVal.(int64); ok {
		md.Set(enum.AuthKey, strconv.Itoa(int(id)))
	}
	requestIdVal := c.Locals(enum.RequestIdKey)
	if requestId, ok := requestIdVal.(string); ok {
		md.Set(enum.RequestIdKey, requestId)
	}
	return metadata.NewOutgoingContext(context.Background(), md)
}

func FromIncoming(ctx context.Context) (int64, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		val := md.Get(enum.AuthKey)
		if len(val) >= 1 {
			id, err := strconv.Atoi(val[0])
			if err != nil {
				return 0, false
			}
			return int64(id), true
		}
	}
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		val := md.Get(enum.AuthKey)
		if len(val) >= 1 {
			id, err := strconv.Atoi(val[0])
			if err != nil {
				return 0, false
			}
			return int64(id), true
		}
	}
	return 0, false
}

func TraceContext(ctx context.Context) context.Context {
	inMD := metadata.Pairs()
	outMD := metadata.Pairs()
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		inMD = md
	}
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		outMD = md
	}
	return metadata.NewOutgoingContext(ctx, metadata.Join(inMD, outMD))
}

func BuildAuthContext(userId int64) context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{enum.AuthKey: fmt.Sprintf("%d", userId)}))
}

func MockIncomingContext() context.Context {
	return metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{enum.AuthKey: strconv.Itoa(enum.SuperUserID)}))
}

func MockOutgoingContext() context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{enum.AuthKey: strconv.Itoa(enum.SuperUserID)}))
}

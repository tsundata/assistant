package md

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tsundata/assistant/api/enum"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func Outgoing(c *fiber.Ctx) context.Context {
	val := c.Locals(enum.AuthKey)
	md := metadata.New(map[string]string{})
	if id, ok := val.(int64); ok {
		md.Append(enum.AuthKey, strconv.Itoa(int(id)))
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
	return 0, false
}

func MockIncomingContext() context.Context {
	return metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"id": strconv.Itoa(enum.SuperUserID)}))
}

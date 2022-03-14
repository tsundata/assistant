package health

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
	"time"
)

type Client struct {
	Status sync.Map
}

func NewHealthClient(client *rpc.Client) *Client {
	hc := &Client{}
	hc.watch(client, enum.Id)
	hc.watch(client, enum.Chatbot)
	hc.watch(client, enum.Message)
	hc.watch(client, enum.Middle)
	hc.watch(client, enum.Storage)
	hc.watch(client, enum.User)
	return hc
}

func (hc *Client) watch(client *rpc.Client, service string) {
	conn, err := client.Dial(service, rpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	c := grpc_health_v1.NewHealthClient(conn)

	go func() {
		for {
			reply, err := c.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{Service: service})
			if err != nil || reply.Status != grpc_health_v1.HealthCheckResponse_SERVING {
				hc.Status.Store(service, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
			} else {
				hc.Status.Store(service, grpc_health_v1.HealthCheckResponse_SERVING)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}

var ProviderSet = wire.NewSet(NewHealthClient)

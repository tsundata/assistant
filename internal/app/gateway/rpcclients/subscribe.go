package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewSubscribeClient(client *rpc.Client) (pb.SubscribeClient, error) {
	conn, err := client.Dial("subscribe")
	if err != nil {
		return nil, err
	}
	c := pb.NewSubscribeClient(conn)

	return c, nil
}

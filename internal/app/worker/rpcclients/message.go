package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewMessageClient(client *rpc.Client) (pb.MessageClient, error) {
	conn, err := client.Dial("message")
	if err != nil {
		return nil, err
	}
	c := pb.NewMessageClient(conn)

	return c, nil
}

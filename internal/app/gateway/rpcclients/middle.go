package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewMiddleClient(client *rpc.Client) (pb.MiddleClient, error) {
	conn, err := client.Dial("middle")
	if err != nil {
		return nil, err
	}
	c := pb.NewMiddleClient(conn)

	return c, nil
}

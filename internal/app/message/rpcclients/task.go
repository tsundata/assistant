package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewTaskClient(client *rpc.Client) (pb.TaskClient, error) {
	conn, err := client.Dial("task")
	if err != nil {
		return nil, err
	}
	c := pb.NewTaskClient(conn)

	return c, nil
}

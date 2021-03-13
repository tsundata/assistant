package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewWorkflowClient(client *rpc.Client) (pb.WorkflowClient, error) {
	conn, err := client.Dial("workflow")
	if err != nil {
		return nil, err
	}
	c := pb.NewWorkflowClient(conn)

	return c, nil
}

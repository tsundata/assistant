package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewWorkflowClient(client *rpc.Client) (pb.WorkflowSvcClient, error) {
	conn, err := client.Dial(enum.Workflow, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewWorkflowSvcClient(conn)
	return c, nil
}

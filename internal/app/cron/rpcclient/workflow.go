package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewWorkflowClient(client *rpc.Client) (pb.WorkflowClient, error) {
	conn, err := client.Dial(app.Workflow, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewWorkflowClient(conn)
	return c, nil
}

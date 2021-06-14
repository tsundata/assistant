package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	workflowOnce   sync.Once
	workflowClient pb.WorkflowClient
)

func GetWorkflowClient(client *rpc.Client) pb.WorkflowClient {
	workflowOnce.Do(func() {
		conn, err := client.Dial("workflow")
		if err != nil {
			fmt.Println(err)
			return
		}
		workflowClient = pb.NewWorkflowClient(conn)
	})

	return workflowClient
}

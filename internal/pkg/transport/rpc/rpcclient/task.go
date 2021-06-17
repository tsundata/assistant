package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	taskOnce   sync.Once
	taskClient pb.TaskClient
)

func GetTaskClient(client *rpc.Client) pb.TaskClient {
	taskOnce.Do(func() {
		conn, err := client.Dial(app.Task)
		if err != nil {
			fmt.Println(err)
			return
		}
		taskClient = pb.NewTaskClient(conn)
	})

	return taskClient
}

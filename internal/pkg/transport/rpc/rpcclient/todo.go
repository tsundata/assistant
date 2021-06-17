package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	todoOnce   sync.Once
	todoClient pb.TodoClient
)

func GetTodoClient(client *rpc.Client) pb.TodoClient {
	todoOnce.Do(func() {
		conn, err := client.Dial(app.Todo)
		if err != nil {
			fmt.Println(err)
			return
		}
		todoClient = pb.NewTodoClient(conn)
	})

	return todoClient
}

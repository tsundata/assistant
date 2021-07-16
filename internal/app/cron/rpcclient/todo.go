package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewTodoClient(client *rpc.Client) (pb.TodoSvcClient, error) {
	conn, err := client.Dial(app.Todo, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "todo client dial error")
	}
	c := pb.NewTodoSvcClient(conn)
	return c, nil
}

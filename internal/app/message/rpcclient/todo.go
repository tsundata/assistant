package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewTodoClient(client *rpc.Client) (pb.TodoClient, error) {
	conn, err := client.Dial(app.Todo)
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewTodoClient(conn)
	return c, nil
}

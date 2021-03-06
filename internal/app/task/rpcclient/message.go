package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewMessageClient(client *rpc.Client) (pb.MessageClient, error) {
	conn, err := client.Dial(app.Message, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "message client dial error")
	}
	c := pb.NewMessageClient(conn)
	return c, nil
}

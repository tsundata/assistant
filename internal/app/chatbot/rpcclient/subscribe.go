package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewSubscribe(client *rpc.Client) (pb.SubscribeClient, error) {
	conn, err := client.Dial(app.Subscribe, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "subscribe client dial error")
	}
	c := pb.NewSubscribeClient(conn)
	return c, nil
}

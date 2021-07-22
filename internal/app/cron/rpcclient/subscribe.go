package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewSubscribe(client *rpc.Client) (pb.SubscribeSvcClient, error) {
	conn, err := client.Dial(enum.Subscribe, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "subscribe client dial error")
	}
	c := pb.NewSubscribeSvcClient(conn)
	return c, nil
}

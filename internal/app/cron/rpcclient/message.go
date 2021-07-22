package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewMessageClient(client *rpc.Client) (pb.MessageSvcClient, error) {
	conn, err := client.Dial(enum.Message, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "message client dial error")
	}
	c := pb.NewMessageSvcClient(conn)
	return c, nil
}

package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewMiddleClient(client *rpc.Client) (pb.MiddleSvcClient, error) {
	conn, err := client.Dial(enum.Middle, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewMiddleSvcClient(conn)
	return c, nil
}

package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewIdClient(client *rpc.Client) (pb.IdSvcClient, error) {
	conn, err := client.Dial(enum.Id, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "id client dial error")
	}
	c := pb.NewIdSvcClient(conn)
	return c, nil
}

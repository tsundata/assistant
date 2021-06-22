package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewMiddleClient(client *rpc.Client) (pb.MiddleClient, error) {
	conn, err := client.Dial(app.Middle)
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewMiddleClient(conn)
	return c, nil
}

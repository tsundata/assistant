package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewUserClient(client *rpc.Client) (pb.UserClient, error) {
	conn, err := client.Dial(app.Middle)
	if err != nil {
		return nil, errors.Wrap(err, "middle client dial error")
	}
	c := pb.NewUserClient(conn)
	return c, nil
}

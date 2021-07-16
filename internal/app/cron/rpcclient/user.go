package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewUserClient(client *rpc.Client) (pb.UserSvcClient, error) {
	conn, err := client.Dial(app.User, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewUserSvcClient(conn)
	return c, nil
}

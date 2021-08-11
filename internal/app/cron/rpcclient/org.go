package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewOrgClient(client *rpc.Client) (pb.OrgSvcClient, error) {
	conn, err := client.Dial(enum.Org, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewOrgSvcClient(conn)
	return c, nil
}

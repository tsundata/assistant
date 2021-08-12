package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewFinanceClient(client *rpc.Client) (pb.FinanceSvcClient, error) {
	conn, err := client.Dial(enum.Finance, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewFinanceSvcClient(conn)
	return c, nil
}

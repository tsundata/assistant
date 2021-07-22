package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewNLPClient(client *rpc.Client) (pb.NLPSvcClient, error) {
	conn, err := client.Dial(enum.NLP, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "nlp client dial error")
	}
	c := pb.NewNLPSvcClient(conn)
	return c, nil
}

package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewNLPClient(client *rpc.Client) (pb.NLPClient, error) {
	conn, err := client.Dial(app.NLP)
	if err != nil {
		return nil, errors.Wrap(err, "nlp client dial error")
	}
	c := pb.NewNLPClient(conn)
	return c, nil
}

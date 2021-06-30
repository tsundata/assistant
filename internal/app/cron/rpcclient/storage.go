package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewStorageClient(client *rpc.Client) (pb.StorageClient, error) {
	conn, err := client.Dial(app.Storage)
	if err != nil {
		return nil, errors.Wrap(err, "storage client dial error")
	}
	c := pb.NewStorageClient(conn)
	return c, nil
}

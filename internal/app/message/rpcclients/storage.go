package rpcclients

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func NewStorageClient(client *rpc.Client) (pb.StorageClient, error) {
	conn, err := client.Dial("storage")
	if err != nil {
		return nil, err
	}
	c := pb.NewStorageClient(conn)

	return c, nil
}

package rpcclient

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewMessageClient(_ *rpc.Client) (pb.MessageClient, error) {
	return nil, nil
}

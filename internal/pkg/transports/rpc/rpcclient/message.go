package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	messageOnce   sync.Once
	messageClient pb.MessageClient
)

func GetMessageClient(client *rpc.Client) pb.MessageClient {
	messageOnce.Do(func() {
		conn, err := client.Dial("message")
		if err != nil {
			fmt.Println(err)
			return
		}
		messageClient = pb.NewMessageClient(conn)
	})

	return messageClient
}

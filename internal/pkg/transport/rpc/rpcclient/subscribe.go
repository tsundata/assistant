package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	subscribeOnce   sync.Once
	subscribeClient pb.SubscribeClient
)

func GetSubscribeClient(client *rpc.Client) pb.SubscribeClient {
	subscribeOnce.Do(func() {
		conn, err := client.Dial(app.Subscribe)
		if err != nil {
			fmt.Println(err)
			return
		}
		subscribeClient = pb.NewSubscribeClient(conn)
	})

	return subscribeClient
}

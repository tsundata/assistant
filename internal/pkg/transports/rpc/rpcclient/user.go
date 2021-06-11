package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	userOnce   sync.Once
	userClient pb.UserClient
)

func GetUserClient(client *rpc.Client) pb.UserClient {
	userOnce.Do(func() {
		conn, err := client.Dial("user")
		if err != nil {
			fmt.Println(err)
			return
		}
		userClient = pb.NewUserClient(conn)
	})

	return userClient
}

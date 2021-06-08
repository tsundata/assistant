package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	middleOnce   sync.Once
	middleClient pb.MiddleClient
)

func GetMiddleClient(client *rpc.Client) pb.MiddleClient {
	middleOnce.Do(func() {
		conn, err := client.Dial("middle")
		if err != nil {
			fmt.Println(err)
			return
		}
		middleClient = pb.NewMiddleClient(conn)
	})

	return middleClient
}

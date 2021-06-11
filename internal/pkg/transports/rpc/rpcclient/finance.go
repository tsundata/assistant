package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	financeOnce   sync.Once
	financeClient pb.FinanceClient
)

func GetFinanceClient(client *rpc.Client) pb.FinanceClient {
	financeOnce.Do(func() {
		conn, err := client.Dial("finance")
		if err != nil {
			fmt.Println(err)
			return
		}
		financeClient = pb.NewFinanceClient(conn)
	})

	return financeClient
}

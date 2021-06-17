package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"sync"
)

var (
	financeOnce   sync.Once
	financeClient pb.FinanceClient
)

func GetFinanceClient(client *rpc.Client) pb.FinanceClient {
	financeOnce.Do(func() {
		conn, err := client.Dial(app.Finance)
		if err != nil {
			fmt.Println(err)
			return
		}
		financeClient = pb.NewFinanceClient(conn)
	})

	return financeClient
}
